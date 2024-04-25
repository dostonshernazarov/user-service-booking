package repository

import (
	"context"
	"Booking/user-service-booking/internal/entity"
)

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, params map[string]string) (*entity.User, error)
	ListUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
	ListDeletedUsers(ctx context.Context, limit, offset int64) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	SoftDelete(ctx context.Context, id string) error
	HardDelete(ctx context.Context, id string) error
	UserEstablishment(ctx context.Context, id, user_id, establishment_id string) (string, error)
}
