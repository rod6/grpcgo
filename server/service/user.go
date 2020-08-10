package service

import "demo/pkg/user"

type UserService struct {
	user.UnimplementedUserServiceServer
}
