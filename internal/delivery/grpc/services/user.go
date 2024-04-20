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

func (s userRPC) Get(ctx context.Context, id *pb.IdReq) (*pb.User, error) {
	return &pb.User{Id: id.Id}, nil
}

func (s userRPC) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
    return &pb.User{
        Id: req.Id,
        FullName: req.FullName,
    }, nil
}

func (s userRPC) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
    return &pb.User{
        Id: req.Id,
        FullName: req.FullName,
    }, nil
}

func (s userRPC) SoftDelete(ctx context.Context, req *pb.IdReq) (*pb.DelRes, error) {
    return &pb.DelRes{}, nil
}

func (s userRPC) List(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
    return &pb.ListUsersRes{}, nil
}

func (s userRPC) HardDelete(ctx context.Context, req *pb.IdReq) (*pb.DelRes, error) {
    return &pb.DelRes{}, nil
}