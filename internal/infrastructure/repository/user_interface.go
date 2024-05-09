package repository

import (
	"context"
	"Booking/user-service-booking/internal/entity"
)

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, params map[string]string) (*entity.User, error)
	ListUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error)
	ListDeletedUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	SoftDelete(ctx context.Context, id string) error

	UserEstablishmentCreate(ctx context.Context, user_id, establishment_id string) (string, string, error)
	UserEstablishmentGet(ctx context.Context, params map[string]string) (*entity.User, string, error)
	UserEstablishmentDelete(ctx context.Context, params map[string]string) error

	CheckUniquess(ctx context.Context, field, value string) (int32, error)
	Exists(ctx context.Context, field, value string) (*entity.User, error)
}
