package services

import (
	pb "Booking/user-service-booking/genproto/user-proto"
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/pkg/otlp"
	"Booking/user-service-booking/internal/usecase"
	"Booking/user-service-booking/internal/usecase/event"
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type userRPC struct {
	logger         *zap.Logger
	userUsecase    usecase.User
	brokerProducer event.BrokerProducer
}

func NewRPC(logger *zap.Logger, userUsecase usecase.User, brokerProducer event.BrokerProducer) pb.UserServiceServer {
	return &userRPC{
		logger:         logger,
		userUsecase:    userUsecase,
		brokerProducer: brokerProducer,
	}
}

func (s userRPC) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
    ctx, span := otlp.Start(ctx, "Delivery", "Create")
    span.SetAttributes(
      attribute.Key("deliveryCreateId").String(req.Id),
    )
    defer span.End()
	
	createdUser, err := s.userUsecase.Create(ctx, &entity.User{
		Id:           req.Id,
		FullName:     req.FullName,
		Email:        req.Email,
		Password:     req.Password,
		DateOfBirth:  req.DateOfBirth,
		ProfileImg:   req.ProfileImg,
		Card:         req.Card,
		Gender:       req.Gender,
		PhoneNumber:  req.PhoneNumber,
		Role:         req.Role,
		RefreshToken: req.RefreshToken,
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
    ctx, span := otlp.Start(ctx, "Delivery", "Get")
    span.SetAttributes(
      attribute.Key("deliveryGetId").String(filter.Filter["id"]),
    )
    defer span.End()
	
	filterUser, err := s.userUsecase.Get(ctx, filter.Filter)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetUser{
		User: &pb.User{
			Id: 			filterUser.Id,
			FullName:		filterUser.FullName,
			Email:			filterUser.Email,
			Password: filterUser.Password,
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
	ctx, span := otlp.Start(ctx, "Delivery", "ListUsers")
    span.SetAttributes(
      attribute.Key("deliveryListLimit").String(fmt.Sprint(req.Limit)),
      attribute.Key("deliveryListOffset").String(fmt.Sprint(req.Offset)),
    )
    defer span.End()
	
	listedUsers, count, err := s.userUsecase.ListUsers(ctx, int64(req.Limit), int64(req.Offset), req.Fv.Field, req.Fv.Value)
	if err != nil {
		return nil, err
	}
	var users []*pb.UserList
	for _, user := range listedUsers {
		users = append(users, &pb.UserList{
			Id:           user.Id,
			FullName:     user.FullName,
			Email:        user.Email,
			DateOfBirth:  user.DateOfBirth,
			ProfileImg:   user.ProfileImg,
			Card:         user.Card,
			Gender:       user.Gender,
			PhoneNumber:  user.PhoneNumber,
			Role:         user.Role,
			RefreshToken: user.RefreshToken,
			CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}
	return &pb.ListUsersRes{
		Users: users,
		Count: count,
	}, nil
}

func (s userRPC) ListDeletedUsers(ctx context.Context, req *pb.ListUsersReq) (*pb.ListUsersRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "ListDeletedUsers")
    span.SetAttributes(
      attribute.Key("deliveryListDeletedLimit").String(fmt.Sprint(req.Limit)),
      attribute.Key("deliveryListDeletedOffset").String(fmt.Sprint(req.Offset)),
    )
    defer span.End()
	
	gotAllUsers, count, err := s.userUsecase.ListDeletedUsers(ctx, int64(req.Limit), int64(req.Offset), req.Fv.Field, req.Fv.Value)
	if err != nil {
		return nil, err
	}
	var users []*pb.UserList
	for _, user := range gotAllUsers {
		users = append(users, &pb.UserList{
			Id:           user.Id,
			FullName:     user.FullName,
			Email:        user.Email,
			DateOfBirth:  user.DateOfBirth,
			ProfileImg:   user.ProfileImg,
			Card:         user.Card,
			Gender:       user.Gender,
			PhoneNumber:  user.PhoneNumber,
			Role:         user.Role,
			RefreshToken: user.RefreshToken,
			CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			DeletedAt:    user.DeletedAt.Format("2006-01-02T15:04:05Z"),
		})
	}
	return &pb.ListUsersRes{
		Users: users,
		Count: count,
	}, nil
}

func (s userRPC) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "Update")
    span.SetAttributes(
      attribute.Key("deliveryUpdateId").String(fmt.Sprint(req.Id)),
    )
    defer span.End()
	
	updatedUser, err := s.userUsecase.Update(ctx, &entity.User{
		Id:           req.Id,
		FullName:     req.FullName,
		Email:        req.Email,
		Password:     req.Password,
		DateOfBirth:  req.DateOfBirth,
		ProfileImg:   req.ProfileImg,
		Card:         req.Card,
		Gender:       req.Gender,
		PhoneNumber:  req.PhoneNumber,
		Role:         req.Role,
		RefreshToken: req.RefreshToken,
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
	ctx, span := otlp.Start(ctx, "Delivery", "Delete")
    span.SetAttributes(
      attribute.Key("deliveryDeleteId").String(fmt.Sprint(req.Id)),
    )
    defer span.End()
	
	err := s.userUsecase.SoftDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{}, nil
}

func (s userRPC) UserEstablishmentCreate(ctx context.Context, req *pb.UE) (*pb.UE, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UECreate")
    span.SetAttributes(
      attribute.Key("deliveryUECreateUId").String(fmt.Sprint(req.UserId)),
      attribute.Key("deliveryUECreateEId").String(fmt.Sprint(req.EstablishmentId)),
    )
    defer span.End()
	
	user_id, establishment_id, err := s.userUsecase.UserEstablishmentCreate(ctx, req.UserId, req.EstablishmentId)
	if err != nil {
		return nil, err
	}
	return &pb.UE{
		UserId:          user_id,
		EstablishmentId: establishment_id,
	}, nil
}

func (s userRPC) UserEstablishmentGet(ctx context.Context, req *pb.Filter) (*pb.UEwU, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UEGet")
    span.SetAttributes(
      attribute.Key("deliveryUEGetUId").String(fmt.Sprint(req.Filter["user_id"])),
      attribute.Key("deliveryUEGetEId").String(fmt.Sprint(req.Filter["establishment_id"])),
    )
    defer span.End()
	
	user, establishment_id, err := s.userUsecase.UserEstablishmentGet(ctx, req.Filter)
	if err != nil {
		return nil, err
	}
	return &pb.UEwU{
		User: &pb.User{
			Id:           user.Id,
			FullName:     user.FullName,
			Email:        user.Email,
			Password:     user.Password,
			DateOfBirth:  user.DateOfBirth,
			ProfileImg:   user.ProfileImg,
			Card:         user.Card,
			Gender:       user.Gender,
			PhoneNumber:  user.PhoneNumber,
			Role:         user.Role,
			RefreshToken: user.RefreshToken,
			CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
		EstablishmentId: establishment_id,
	}, nil
}

func (s userRPC) UserEstablishmentDelete(ctx context.Context, req *pb.Filter) (*pb.DelRes, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "UEDelete")
    span.SetAttributes(
      attribute.Key("deliveryUEDeleteUId").String(fmt.Sprint(req.Filter["user_id"])),
      attribute.Key("deliveryUEDeleteEId").String(fmt.Sprint(req.Filter["establishment_id"])),
    )
    defer span.End()
	
	err := s.userUsecase.UserEstablishmentDelete(ctx, req.Filter)
	if err != nil {
		return nil, err
	}
	return &pb.DelRes{}, nil
}

func (s userRPC) CheckUniquess(ctx context.Context, req *pb.FV) (*pb.Status, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "CheckUniquess")
    span.SetAttributes(
      attribute.Key("deliveryCheckUniquessField").String(fmt.Sprint(req.Field)),
      attribute.Key("deliveryCheckUniquessValue").String(fmt.Sprint(req.Value)),
    )
    defer span.End()
    
    status, err := s.userUsecase.CheckUniquess(ctx, req.Field, req.Value)
    if err!= nil {
        return nil, err
    }
    return &pb.Status{
        Code: status,
    }, nil
}

func (s userRPC) Exists(ctx context.Context, req *pb.FV) (*pb.User, error) {
	ctx, span := otlp.Start(ctx, "Delivery", "Exists")
	span.SetAttributes(
      attribute.Key("deliveryExistsField").String(fmt.Sprint(req.Field)),
      attribute.Key("deliveryExistsValue").String(fmt.Sprint(req.Value)),
    )
	defer span.End()

	user, err := s.userUsecase.Exists(ctx, req.Field, req.Value)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:             user.Id,
        FullName:       user.FullName,
        Email:          user.Email,
        Password:        user.Password,
        DateOfBirth:    user.DateOfBirth,
        ProfileImg:     user.ProfileImg,
        Card:           user.Card,
        Gender:         user.Gender,
        PhoneNumber:    user.PhoneNumber,
        Role:           user.Role,
        RefreshToken:   user.RefreshToken,
        CreatedAt:      user.CreatedAt.Format("2006-01-02T15:04:05Z"),
        UpdatedAt:      user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
    }, nil
}
