package player

import (
	"context"
	"fmt"
	"log"
	pb "matchmaking/cmd/proto"
	"matchmaking/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PlayerServiceServer struct {
	pb.UnimplementedPlayerServiceServer
	repo PlayerRepo
	cfg  config.Config
}

func NewPlayerServiceServer(repo PlayerRepo, cfg config.Config) *PlayerServiceServer {
	return &PlayerServiceServer{repo: repo, cfg: cfg}
}

func (s *PlayerServiceServer) Create(ctx context.Context, req *pb.PlayerCreateRequest) (*pb.PlayerCreateResponse, error) {
	if req.Username == "" || req.Email == "" || len(req.Password) == 0 {
		return nil, status.Error(codes.InvalidArgument, "username, email, and password required")
	}

	exists, _ := s.repo.Get(ctx, req.Username)
	if exists.Username != "" {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, status.Error(codes.Internal, "password hash failed")
	}

	p := Player{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hash,
	}

	id, err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to create player. please try again.")
	}

	return &pb.PlayerCreateResponse{Id: fmt.Sprintf("%d", id)}, nil
}

func (s *PlayerServiceServer) Login(ctx context.Context, req *pb.PlayerLoginRequest) (*pb.PlayerLoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	if username == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "missing credentials")
	}

	// Fetch the user from the database
	player, err := s.repo.Get(ctx, username)
	if err != nil {
		return nil, status.Error(codes.Internal, "something went wrong. please try again.")
	}

	// Compare the password and stored hash
	err = bcrypt.CompareHashAndPassword(player.HashedPassword, []byte(password))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid credentials.")
	}

	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		// add "iat", "iss", etc. as needed
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(s.cfg.JWTSecret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.Internal, "something went wrong. please try again.")
	}

	return &pb.PlayerLoginResponse{Token: tokenString}, nil
}
