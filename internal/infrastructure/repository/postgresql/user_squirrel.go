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
	userTableName      = "user"
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
	return p.db.Sq.Builder.
		Select(
			"id",
			"first_name",
			"last_name",
		).From(p.tableName)
}

func (p userRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Create")
	// defer span.End()

	data := map[string]any{
		"id":			user.Id,
		"full_name":	user.FullName,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return user, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return user, p.db.Error(err)
	}

	return user, nil
}

func (p userRepo) Get(ctx context.Context, id string) (*entity.User, error) {
	var (
		user entity.User
	)

	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Get")
	// defer span.End()

	queryBuilder := p.userSelectQueryPrefix()

	queryBuilder = queryBuilder.Where(p.db.Sq.Equal("id", id))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "get"))
	}
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&user.Id,
		&user.FullName,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &user, nil
}

func (p userRepo) List(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	var (
		users []*entity.User
	)
	queryBuilder := p.userSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()
	users = make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
			&user.Id,
			&user.FullName,
		); err != nil {
			return nil, p.db.Error(err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p userRepo) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	clauses := map[string]any{
		"full_name":	user.FullName,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", user.Id)).
		ToSql()
	if err != nil {
		return user, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return user, p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return user, p.db.Error(fmt.Errorf("no sql rows"))
	}

	return user, nil
}

func (p userRepo) HardDelete(ctx context.Context, id string) error {
	sqlStr, args, err := p.db.Sq.Builder.Delete(p.tableName).Where(p.db.Sq.Equal("id", id)).ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (p userRepo) SoftDelete(ctx context.Context, id string) error {
	clauses := map[string]any{
		"deleted_at":	time.Now().UTC(),
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", id)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}