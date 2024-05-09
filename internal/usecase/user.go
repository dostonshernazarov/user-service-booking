package usecase

import (
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/infrastructure/repository"
	"Booking/user-service-booking/internal/pkg/otlp"
	"context"
	"time"
)

const (
	serviceNameUser = "userService"
	spanNameUser    = "userUsecase"
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

type UserService struct {
	BaseUseCase
	repo       repository.User
	ctxTimeout time.Duration
}

func NewUserService(ctxTimeout time.Duration, repo repository.User) UserService {
	return UserService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u UserService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Create")
	defer span.End()

	u.beforeRequest(nil, &user.CreatedAt, &user.UpdatedAt, nil)

	return u.repo.Create(ctx, user)
}

func (u UserService) Get(ctx context.Context, params map[string]string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Get")
	defer span.End()

	return u.repo.Get(ctx, params)
}

func (u UserService) ListUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"ListUsers")
	defer span.End()

	return u.repo.ListUsers(ctx, limit, offset, field, value)
}

func (u UserService) ListDeletedUsers(ctx context.Context, limit, offset int64, field, value string) ([]*entity.User, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"ListDeletedUsers")
	defer span.End()

	return u.repo.ListDeletedUsers(ctx, limit, offset, field, value)
}

func (u UserService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Update")
	defer span.End()

	u.beforeRequest(nil, nil, &user.UpdatedAt, nil)

	return u.repo.Update(ctx, user)
}

func (u UserService) SoftDelete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Delete")
	defer span.End()
	
	var user entity.User
	user.Id = id
    user.DeletedAt = time.Now().UTC()
	u.beforeRequest(nil, nil, nil, &user.DeletedAt)

	return u.repo.SoftDelete(ctx, id)
}

func (u UserService) UserEstablishmentCreate(ctx context.Context, user_id, establishment_id string) (string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"UECreate")
	defer span.End()

    return u.repo.UserEstablishmentCreate(ctx, user_id, establishment_id)
}

func (u UserService) UserEstablishmentGet(ctx context.Context, params map[string]string) (*entity.User, string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"UEGet")
	defer span.End()

    return u.repo.UserEstablishmentGet(ctx, params)
}

func (u UserService) UserEstablishmentDelete(ctx context.Context, params map[string]string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"UEDelete")
	defer span.End()

    return u.repo.UserEstablishmentDelete(ctx, params)
}

func (u UserService) CheckUniquess(ctx context.Context, field, value string) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"CheckUniquess")
    defer span.End()

    return u.repo.CheckUniquess(ctx, field, value)
}

func (u UserService) Exists(ctx context.Context, field, value string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Exists")
    defer span.End()

    return u.repo.Exists(ctx, field, value)
}