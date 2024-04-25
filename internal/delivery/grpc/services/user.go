package services

import (
	pb "Booking/user-service-booking/genproto/user-proto"
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/usecase"
	"Booking/user-service-booking/internal/usecase/event"
	"context"

	"go.uber.org/zap"
)

type userRPC struct {
	logger                *zap.Logger
	userUsecase           usecase.User
	brokerProducer        event.BrokerProducer
}

func NewRPC(logger *zap.Logger, userUsecase usecase.User, brokerProducer event.BrokerProducer) pb.UserServiceServer {
	return &userRPC{
		logger:                logger,
		userUsecase:           userUsecase,
		brokerProducer:        brokerProducer,
	}
}

func (s userRPC) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	createdUser, err := s.userUsecase.Create(ctx, &entity.User{
		Id:             req.Id,
		FullName:       req.FullName,
		Email:          req.Email,
		Password:       req.Password,
		DateOfBirth:	req.DateOfBirth,
		ProfileImg:     req.ProfileImg,
		Card:           req.Card,
		Gender:         req.Gender,
		PhoneNumber:    req.PhoneNumber,
		Role:           req.Role,
		RefreshToken:   req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
    return &pb.User{
		Id:             createdUser.Id,
		FullName:       createdUser.FullName,
		Email:          createdUser.Email,
		Password:       createdUser.Password,
		DateOfBirth:	createdUser.DateOfBirth,
		ProfileImg:     createdUser.ProfileImg,
		Card:           createdUser.Card,
		Gender:         createdUser.Gender,
		PhoneNumber:    createdUser.PhoneNumber,
		Role:           createdUser.Role,
		RefreshToken:   createdUser.RefreshToken,
    }, nil
}

func (s userRPC) Get(ctx context.Context, filter *pb.Filter) (*pb.GetUser, error) {
	filterUser, err := s.userUsecase.Get(ctx, filter.Filter)
	if err!= nil {
        return nil, err
    }

	resp := &pb.GetUser{
		User: &pb.User{
			Id: 			filterUser.Id,
			FullName:		filterUser.FullName,
			Email:			filterUser.Email,
			Password:		filterUser.Password,
			DateOfBirth:	filterUser.DateOfBirth,
			ProfileImg:		filterUser.ProfileImg,
			Card:			filterUser.Card,
			Gender:			filterUser.Gender,
			PhoneNumber:	filterUser.PhoneNumber,
			Role:			filterUser.Role,
			RefreshToken:	filterUser.RefreshToken,
			CreatedAt:		filterUser.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:		filterUser.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}

	return resp, nil
}

func (s userRPC) ListUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
	listedUsers, err := s.userUsecase.ListUsers(ctx, int64(req.Limit), int64(req.Offset))
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for _, user := range listedUsers {
		users = append(users, &pb.User{
            Id:             user.Id,
            FullName:       user.FullName,
            Email:          user.Email,
            Password:		user.Password,
			DateOfBirth:    user.DateOfBirth,
            ProfileImg:     user.ProfileImg,
            Card:           user.Card,
            Gender:         user.Gender,
            PhoneNumber:    user.PhoneNumber,
			Role:           user.Role,
            RefreshToken:   user.RefreshToken,
            CreatedAt:      user.CreatedAt.Format("2006-01-02T15:04:05Z"),
            UpdatedAt:      user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},)
	}
    return &pb.ListUsersRes{
		Users: users,
	}, nil
}

func (s userRPC) ListDeletedUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
	gotAllUsers, err := s.userUsecase.ListDeletedUsers(ctx, int64(req.Limit), int64(req.Offset))
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	for _, user := range gotAllUsers {
		users = append(users, &pb.User{
            Id:             user.Id,
            FullName:       user.FullName,
            Email:          user.Email,
            Password:		user.Password,
			DateOfBirth:    user.DateOfBirth,
            ProfileImg:     user.ProfileImg,
            Card:           user.Card,
            Gender:         user.Gender,
            PhoneNumber:    user.PhoneNumber,
			Role:           user.Role,
            RefreshToken:   user.RefreshToken,
            CreatedAt:      user.CreatedAt.Format("2006-01-02T15:04:05Z"),
            UpdatedAt:      user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			DeletedAt: 		user.DeletedAt.Format("2006-01-02T15:04:05Z"),
		},)
	}
	return &pb.ListUsersRes{
		Users: users,
	}, nil
}

func (s userRPC) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	updatedUser, err := s.userUsecase.Update(ctx, &entity.User{
		Id:             req.Id,
		FullName:       req.FullName,
		Email:          req.Email,
		Password:       req.Password,
		DateOfBirth:	req.DateOfBirth,
		ProfileImg:     req.ProfileImg,
		Card:           req.Card,
		Gender:         req.Gender,
		PhoneNumber:    req.PhoneNumber,
		Role:           req.Role,
		RefreshToken:   req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
    return &pb.User{
		Id:             updatedUser.Id,
		FullName:       updatedUser.FullName,
		Email:          updatedUser.Email,
		Password:       updatedUser.Password,
		DateOfBirth:	updatedUser.DateOfBirth,
		ProfileImg:     updatedUser.ProfileImg,
		Card:           updatedUser.Card,
		Gender:         updatedUser.Gender,
		PhoneNumber:    updatedUser.PhoneNumber,
		Role:           updatedUser.Role,
		RefreshToken:   updatedUser.RefreshToken,
    }, nil
}

func (s userRPC) SoftDelete(ctx context.Context, req *pb.Id) (*pb.DelRes, error) {
	err := s.userUsecase.SoftDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
    return &pb.DelRes{}, nil
}

func (s userRPC) HardDelete(ctx context.Context, req *pb.Id) (*pb.DelRes, error) {
	err := s.userUsecase.HardDelete(ctx, req.Id)
	if err != nil {
        return nil, err
    }
    return &pb.DelRes{}, nil
}

func (s userRPC) UserEstablishmentCreate(ctx context.Context, req *pb.UE) (*pb.UE, error) {
	id, user_id, establishment_id, err := s.userUsecase.UserEstablishmentCreate(ctx, req.Id, req.UserId, req.EstablishmentId)
    if err != nil {
        return nil, err
    }
    return &pb.UE{
		Id: id,
		UserId: user_id,
		EstablishmentId: establishment_id,
	}, nil
}

func (s userRPC) UserEstablishmentGet(ctx context.Context, req *pb.UE) (*pb.UEwU, error) {
	id, user, establishment_id, err := s.userUsecase.UserEstablishmentGet(ctx, req.Id)
    if err!= nil {
        return nil, err
    }
    return &pb.UEwU{
		Id: id,
		User: &pb.User{
			Id:				user.Id,
            FullName:		user.FullName,
            Email:			user.Email,
            Password:		user.Password,
            DateOfBirth:	user.DateOfBirth,
            ProfileImg:		user.ProfileImg,
            Card:			user.Card,
            Gender:			user.Gender,
            PhoneNumber:	user.PhoneNumber,
            Role:			user.Role,
            RefreshToken:	user.RefreshToken,
            CreatedAt:		user.CreatedAt.Format("2006-01-02T15:04:05Z"),
            UpdatedAt:		user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
		EstablishmentId: establishment_id,
	}, nil
}

func (s userRPC) UserEstablishmentDelete(ctx context.Context, req *pb.UE) (*pb.DelRes, error) {
	err := s.userUsecase.UserEstablishmentDelete(ctx, req.Id)
    if err!= nil {
        return nil, err
    }
    return &pb.DelRes{}, nil
}