package main

import (
	"log"

	"github.com/kitoyanok66/inote-auth/internal/auth"
	"github.com/kitoyanok66/inote-auth/internal/auth/jwt"
	"github.com/kitoyanok66/inote-auth/internal/config"
	"github.com/kitoyanok66/inote-auth/internal/database"
	"github.com/kitoyanok66/inote-auth/internal/transport/grpc"
	transportgrpc "github.com/kitoyanok66/inote-auth/internal/transport/grpc"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := auth.NewAuthRepository(db)
	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, cfg.JWTDuration)
	service := auth.NewAuthService(repo, jwtManager)
	handler := grpc.NewAuthHandler(service)

	if err := transportgrpc.RunGRPC(handler, "50051"); err != nil {
		log.Fatalf("Auth gRPC server error: %v", err)
	}
}
