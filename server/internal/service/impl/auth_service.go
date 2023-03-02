package service

import (
	"codeview/config"
	"codeview/internal/dto/request"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/service"
	"codeview/internal/util"
	"context"
	"errors"
	"log"
	"strconv"
	"time"
)

type authService struct {
	cfg        config.AppConfig
	userRepo   repository.UserRepository
	jwtManager *util.JWTManager
}

func NewAuthService(cfg config.AppConfig, userRepo repository.UserRepository) service.AuthService {
	jwtManager := util.NewJWTManager(cfg)

	return &authService{
		cfg,
		userRepo,
		jwtManager,
	}
}

func (s *authService) Login(ctx context.Context, req *request.Login) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("[ERROR] User Service - Login : %v\n", err)
		return "", err
	}

	if ok := util.ComparePasswordHash(req.Password, user.Password); !ok {
		log.Printf("[ERROR] User Service - Login : %v\n", err)
		return "", errors.New("invalid email or password")
	}

	token, err := s.jwtManager.GenerateJwt(strconv.FormatUint(uint64(user.ID), 10), 10*time.Hour)
	if err != nil {
		log.Printf("[ERROR] User Service - Login : %v\n", err)
		return "", err
	}

	return token, nil
}

func (s *authService) Register(ctx context.Context, req *request.Register) (string, error) {
	user, err := s.userRepo.CreateUser(ctx, &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("[ERROR] User Service - Register : %v\n", err)
		return "", err
	}

	token, err := s.jwtManager.GenerateJwt(strconv.FormatUint(uint64(user.ID), 10), 10*time.Hour)
	if err != nil {
		log.Printf("[ERROR] User Service - Register : %v\n", err)
		return "", err
	}

	return token, nil
}
