package postgresql

import (
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/pkg/otlp"
	"Booking/user-service-booking/internal/pkg/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	userTableName              = "users"
	userEstablishmentTableName = "users_establishment"
	userServiceName            = "userService"
	userSpanRepoPrefix         = "userRepo"
)

type userRepo struct {
	userTableName              string
	userEstablishmentTableName string
	db                         *postgres.PostgresDB
}

func NewUserRepo(db *postgres.PostgresDB) *userRepo {
	return &userRepo{
		userTableName:              userTableName,
		userEstablishmentTableName: userEstablishmentTableName,
		db:                         db,
	}
}

func (p *userRepo) userSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"id",
		"full_name",
		"email",
		"password",
		"TO_CHAR(date_of_birth, 'YYYY-MM-DD') AS date_of_birth",
		"profile_img",
		"card",
		"gender",
		"phone_number",
		"role",
		"refresh_token",
		"created_at",
		"updated_at",
	).From(p.userTableName)
}

func (p *userRepo) userSelectQueryPrefixCount() squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"COUNT(*) AS count",
	).From(p.userTableName)
}

func (p *userRepo) userSelectQueryPrefixAdmin() squirrel.SelectBuilder {
	return p.db.Sq.Builder.Select(
		"id",
		"full_name",
		"email",
		"password",
		"TO_CHAR(date_of_birth, 'YYYY-MM-DD') AS date_of_birth",
		"profile_img",
		"card",
		"gender",
		"phone_number",
		"role",
		"refresh_token",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(p.userTableName)
}

func (p userRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Create")
	defer span.End()

	var DOB time.Time
	var err error
	
	if user.DateOfBirth != "" {
		DOB, err = time.Parse("2006-01-02", user.DateOfBirth)

		if err != nil {
			return nil, fmt.Errorf("failed to parse date of birth: %v", err)
		}
	}

	data := map[string]interface{}{
		"id":              user.Id,
		"full_name":       user.FullName,
		"email":           user.Email,
		"password":        user.Password,
		"date_of_birth":   DOB,
		"profile_img":     user.ProfileImg,
		"card":            user.Card,
		"gender":          user.Gender,
		"phone_number":    user.PhoneNumber,
		"role":            user.Role,
		"refresh_token":   user.RefreshToken,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.userTableName).SetMap(data).ToSql()
	if err != nil {
		return user, fmt.Errorf("failed to build SQL query for creating user: %v", err)
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return user, fmt.Errorf("failed to execute SQL query for creating user: %v", err)
	}

	return user, nil
}

func (p userRepo) Get(ctx context.Context, params map[string]string) (*entity.User, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Get")
	defer span.End()

	var user entity.User

	queryBuilder := p.userSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
		if key == "email" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
		if key == "refresh_token" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}

    	queryBuilder = queryBuilder.Where(p.db.Sq.Equal("deleted_at", nil))
	}

	

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query for getting user: %v", err)
	}

	// println("\n\n", query, "\n\n")
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.DateOfBirth,
		&user.ProfileImg,
		&user.Card,
		&user.Gender,
		&user.PhoneNumber,
		&user.Role,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (p userRepo) ListUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"ListUsers")
	defer span.End()

	var (
		users []*entity.User
		count int64
	)

	queryBuilder := p.userSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	}

	queryBuilder = queryBuilder.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.ILike(field, "%"+value+"%"))
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
	}
	defer rows.Close()
	users = make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.ProfileImg,
			&user.Card,
			&user.Gender,
			&user.PhoneNumber,
			&user.Role,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing users: %v", err)
		}
		users = append(users, &user)
	}

	queryCount := p.userSelectQueryPrefixCount()
	query, args, err = queryCount.Where(p.db.Sq.Equal("deleted_at", nil)).Where(p.db.Sq.ILike(field, "%"+value+"%")).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting users: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err!= nil {
        return nil, 0, fmt.Errorf("failed to scan row while counting users: %v", err)
    }

	return users, count, nil
}

func (p userRepo) ListDeletedUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"ListDeletedUsers")
	defer span.End()

	var (
		users []*entity.User
		count int64
	)

	queryBuilder := p.userSelectQueryPrefixAdmin()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	}

	queryBuilder = queryBuilder.Where(p.db.Sq.NotEqual("deleted_at", nil)).Where(p.db.Sq.ILike(field, "%"+value+"%"))
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for listing all users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute SQL query for listing all users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Email,
			&user.Password,
			&user.DateOfBirth,
			&user.ProfileImg,
			&user.Card,
			&user.Gender,
			&user.PhoneNumber,
			&user.Role,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row while listing all users: %v", err)
		}
		users = append(users, &user)
	}

	queryCount := p.userSelectQueryPrefixCount()
	query, args, err = queryCount.Where(p.db.Sq.NotEqual("deleted_at", nil)).Where(p.db.Sq.ILike(field, "%"+value+"%")).ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build SQL query for counting users: %v", err)
	}
	row := p.db.QueryRow(ctx, query, args...)
	if err = row.Scan(&count); err!= nil {
        return nil, 0, fmt.Errorf("failed to scan row while counting users: %v", err)
    }

	return users, count, nil
}

func (p userRepo) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Update")
	defer span.End()

	DOB, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date of birth: %v", err)
	}
	clauses := map[string]interface{}{
		"full_name":     user.FullName,
		"email":         user.Email,
		"password":      user.Password,
		"date_of_birth": DOB,
		"profile_img":   user.ProfileImg,
		"card":          user.Card,
		"gender":        user.Gender,
		"phone_number":  user.PhoneNumber,
		"refresh_token": user.RefreshToken,
	}
	sqlStr, args, err := p.db.Sq.Builder.Update(p.userTableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", user.Id)).
		Where(p.db.Sq.Equal("deleted_at", nil)).
		ToSql()
	if err != nil {
		return user, fmt.Errorf("failed to build SQL query for updating user: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return user, fmt.Errorf("failed to execute SQL query for updating user: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return user, fmt.Errorf("no rows affected while updating user")
	}

	return user, nil
}

func (p userRepo) SoftDelete(ctx context.Context, id string) error {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Delete")
	defer span.End()

	var deletedAt sql.NullTime
	err := p.db.QueryRow(ctx, "SELECT deleted_at FROM "+p.userTableName+" WHERE id = $1", id).Scan(&deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%s: not found", id)
		}
		return fmt.Errorf("failed to query user: %v", err)
	}
	if deletedAt.Valid && !deletedAt.Time.IsZero() {
		return fmt.Errorf("%s: is already soft-deleted", id)
	}

	clauses := map[string]interface{}{
		"deleted_at": time.Now().Format("2006-01-02T15:04:05"),
	}
	sqlBuilder := p.db.Sq.Builder.Update(p.userTableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", id))

	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query for soft deleting user: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query for soft deleting user: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected while soft deleting user")
	}

	return nil
}

func (p userRepo) UserEstablishmentCreate(ctx context.Context, user_id, establishment_id string) (string, string, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"UECreate")
	defer span.End()

	sqlStr, args, err := p.db.Sq.Builder.Insert(p.userEstablishmentTableName).
		Columns("user_id", "establishment_id").
		Values(user_id, establishment_id).
		ToSql()
	if err != nil {
		return "", "", fmt.Errorf("failed to build SQL query for user establishment: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return "", "", fmt.Errorf("failed to execute SQL query for user establishment: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return "", "", fmt.Errorf("no rows affected while user establishment")
	}

	return user_id, establishment_id, nil
}

func (p userRepo) UserEstablishmentGet(ctx context.Context, params map[string]string) (*entity.User, string, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"UEGet")
	defer span.End()

	var user entity.User
	var establishment_id = params["establishment_id"]
	queryBuilder := p.userSelectQueryPrefix()
	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
		queryBuilder = queryBuilder.Where(p.db.Sq.Equal("deleted_at", nil))
	}
    sqlStr, args, err := queryBuilder.
        Join(p.userEstablishmentTableName + " ON " + p.userTableName + ".id = " + p.userEstablishmentTableName + ".user_id").
        // Where(p.db.Sq.Equal("users_establishment.user_id", user_id)).
		Where(p.db.Sq.Equal("users_establishment.establishment_id", establishment_id)).
		ToSql()
	if err != nil {
		return &user, "", fmt.Errorf("failed to build SQL query for user establishment: %v", err)
	}

	rows, err := p.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return &user, "", fmt.Errorf("failed to execute SQL query for user establishment: %v", err)
	}
	defer rows.Close()
	found := false
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Email,
			&user.DateOfBirth,
			&user.ProfileImg,
			&user.Card,
			&user.Gender,
			&user.PhoneNumber,
			&user.Role,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return &user, "", fmt.Errorf("failed to scan row while user establishment: %v", err)
		}
		found = true
	}
	if !found {
		return &user, "", sql.ErrNoRows
	}
	return &user, establishment_id, nil
}

func (p userRepo) UserEstablishmentDelete(ctx context.Context, params map[string]string) error {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"UEDelete")
	defer span.End()

	sqlStr, args, err := p.db.Sq.Builder.Delete(p.userEstablishmentTableName).
        Where(p.db.Sq.Equal("user_id", params["user_id"]), p.db.Sq.Equal("establishment_id", params["establishment_id"])).
        ToSql()
    if err!= nil {
        return fmt.Errorf("failed to build SQL query for user establishment: %v", err)
    }

    commandTag, err := p.db.Exec(ctx, sqlStr, args...)
    if err!= nil {
        return fmt.Errorf("failed to execute SQL query for user establishment: %v", err)
    }

    if commandTag.RowsAffected() == 0 {
        return fmt.Errorf("no rows affected while user establishment")
    }
    
    return nil
}

func (p userRepo) CheckUniquess(ctx context.Context, field, value string) (int32, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"CheckUniquess")
	defer span.End()

	var code int32
	sqlStr, args, err := p.db.Sq.Builder.Select("COUNT(*)").From(p.userTableName).Where(p.db.Sq.Equal(field, value)).Where(p.db.Sq.Equal("deleted_at", nil)).ToSql()
	if err != nil {
		return code, fmt.Errorf("failed to build SQL query for check uniquess: %v", err)
	}
	row := p.db.QueryRow(ctx, sqlStr, args...)
	if err = row.Scan(&code); err != nil {
		return code, fmt.Errorf("failed to scan row while check uniquess: %v", err)
	}
	return code, nil
}

func (p userRepo) Exists(ctx context.Context, field, value string) (*entity.User, error) {
	ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Exists")
	defer span.End()

	var user entity.User
	queryBuilder := p.userSelectQueryPrefix()
	sqlStr, args, err := queryBuilder.Where(p.db.Sq.Equal(field, value)).Where(p.db.Sq.Equal("deleted_at", nil)).ToSql()
	if err != nil {
		return &user, fmt.Errorf("failed to build SQL query for exists: %v", err)
	}
	row := p.db.QueryRow(ctx, sqlStr, args...)
	if err = row.Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.DateOfBirth,
		&user.ProfileImg,
		&user.Card,
		&user.Gender,
		&user.PhoneNumber,
		&user.Role,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan row while checking existence: %v", err)
	}
	return &user, nil
}
