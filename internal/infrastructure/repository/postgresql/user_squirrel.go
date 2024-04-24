package postgresql

import (
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/pkg/postgres"
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	userTableName      = "users"
	userServiceName    = "userService"
	userSpanRepoPrefix = "userRepo"
)

type userRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewUserRepo(db *postgres.PostgresDB) *userRepo {
	return &userRepo{
		tableName: userTableName,
		db:        db,
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
		"establishment_id",
		"refresh_token",
		"created_at",
		"updated_at",
	).From(p.tableName)
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
		"establishment_id",
		"refresh_token",
		"created_at",
		"updated_at",
		"deleted_at",
	).From(p.tableName)
}

func (p userRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	DOB, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date of birth: %v", err)
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
		"establishment_id":user.EstablishmentId,
		"refresh_token":   user.RefreshToken,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
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
	var user entity.User

	queryBuilder := p.userSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
    	queryBuilder = queryBuilder.Where(p.db.Sq.Equal("deleted_at", nil))
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query for getting user: %v", err)
	}
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
		&user.EstablishmentId,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (p userRepo) ListUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	var users []*entity.User

	queryBuilder := p.userSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	}

    queryBuilder = queryBuilder.Where(p.db.Sq.Equal("deleted_at", nil))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query for listing users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query for listing users: %v", err)
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
			&user.EstablishmentId,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row while listing users: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p userRepo) ListDeletedUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	var users []*entity.User

	queryBuilder := p.userSelectQueryPrefixAdmin()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	}

    queryBuilder = queryBuilder.Where(p.db.Sq.NotEqual("deleted_at", nil))
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query for listing all users: %v", err)
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query for listing all users: %v", err)
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
			&user.EstablishmentId,
			&user.RefreshToken,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row while listing all users: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p userRepo) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	DOB, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date of birth: %v", err)
	}
	clauses := map[string]interface{}{
		"full_name":       user.FullName,
		"email":           user.Email,
		"password":        user.Password,
		"date_of_birth":   DOB,
		"profile_img":     user.ProfileImg,
		"card":            user.Card,
		"gender":          user.Gender,
		"phone_number":    user.PhoneNumber,
	}
	sqlStr, args, err := p.db.Sq.Builder.Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", user.Id), p.db.Sq.Equal("deleted_at", nil)).
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
	clauses := map[string]interface{}{
		"deleted_at": time.Now().Format("2006-01-02T15:04:05"),
	}
	sqlBuilder := p.db.Sq.Builder.Update(p.tableName).
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

func (p userRepo) HardDelete(ctx context.Context, id string) error {
	sqlStr, args, err := p.db.Sq.Builder.Delete(p.tableName).Where(p.db.Sq.Equal("id", id)).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query for hard deleting user: %v", err)
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query for hard deleting user: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected while hard deleting user")
	}

	return nil
}
