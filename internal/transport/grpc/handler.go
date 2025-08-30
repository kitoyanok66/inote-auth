package grpc

import (
	"context"

	"github.com/kitoyanok66/inote-auth/internal/auth"
	pb "github.com/kitoyanok66/inote-protos/proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler struct {
	service auth.AuthService
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userDTO, token, err := h.service.Register(ctx, req.Email, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		User: &pb.User{
			Id:        userDTO.ID,
			Email:     userDTO.Email,
			Username:  userDTO.Username,
			CreatedAt: timestamppb.New(userDTO.CreatedAt),
		},
		Token: token,
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userDTO, token, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		User: &pb.User{
			Id:        userDTO.ID,
			Email:     userDTO.Email,
			Username:  userDTO.Username,
			CreatedAt: timestamppb.New(userDTO.CreatedAt),
		},
		Token: token,
	}, nil
}

func (h *AuthHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	userDTO, err := h.service.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:        userDTO.ID,
			Email:     userDTO.Email,
			Username:  userDTO.Username,
			CreatedAt: timestamppb.New(userDTO.CreatedAt),
		},
	}, nil
}
