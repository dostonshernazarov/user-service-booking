package usecase

import (
	"context"
	"time"
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/infrastructure/repository"
)

// const (
// 	serviceNameuser = "userService"
// 	spanNameuser    = "userUsecase"
// )

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context, limit, offset int64) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	SoftDelete(ctx context.Context, id string) error
	HardDelete(ctx context.Context, id string) error
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

	u.beforeRequest(&user.Id, &user.CreatedAt, &user.UpdatedAt, nil)

	return u.repo.Create(ctx, user)
}

func (u UserService) Get(ctx context.Context, id string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Get(ctx, id)
}

func (u UserService) List(ctx context.Context, limit, offset int64) ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.List(ctx, limit, offset)
}

func (u UserService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(nil, nil, &user.UpdatedAt, nil)

	return u.repo.Update(ctx, user)
}

func (u UserService) SoftDelete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()
	
	var user entity.User
	user.Id = id
    user.DeletedAt = time.Now().UTC()
	u.beforeRequest(nil, nil, nil, &user.DeletedAt)

	return u.repo.SoftDelete(ctx, id)
}

func (u UserService) HardDelete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
    defer cancel()

    return u.repo.HardDelete(ctx, id)
}