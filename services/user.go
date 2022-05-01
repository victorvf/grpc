package services

import (
	"context"
	"fmt"

	"github.com/victorvf/grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)

	return &pb.User{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
