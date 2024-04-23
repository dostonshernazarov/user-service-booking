package services

import (
	"context"
	pb "Booking/user-service-booking/genproto/user-proto"
	"Booking/user-service-booking/internal/usecase"
	"Booking/user-service-booking/internal/usecase/event"

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
    return &pb.User{
        Id: req.Id,
        FullName: req.FullName,
    }, nil
}

func (s userRPC) Get(ctx context.Context, filter *pb.Filter) (*pb.GetUser, error) {
	filterUser, err := s.userUsecase.Get(ctx, filter.Filter)
	if err!= nil {
        return nil, err
    }

	resp := &pb.GetUser{
		User: &pb.User{
			Id: filterUser.Id,
			FullName: filterUser.FullName,
			Email: filterUser.Email,
			Password: filterUser.Password,
			DateOfBirth: filterUser.DateOfBirth,
			ProfileImg: filterUser.ProfileImg,
			Card: filterUser.Card,
			Gender: filterUser.Gender,
			PhoneNumber: filterUser.PhoneNumber,
			Role: filterUser.Role,
			EstablishmentId: filterUser.EstablishmentId,
			RefreshToken: filterUser.RefreshToken,
			CreatedAt: filterUser.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: filterUser.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}

	return resp, nil
}

func (s userRPC) ListUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
    return &pb.ListUsersRes{}, nil
}

func (s userRPC) GetAllUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
	return &pb.ListUsersRes{}, nil
}

func (s userRPC) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
    return &pb.User{
        Id: req.Id,
        FullName: req.FullName,
    }, nil
}

func (s userRPC) SoftDelete(ctx context.Context, req *pb.Filter) (*pb.DelRes, error) {
    return &pb.DelRes{}, nil
}

func (s userRPC) HardDelete(ctx context.Context, req *pb.Filter) (*pb.DelRes, error) {
    return &pb.DelRes{}, nil
}