package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewUserRepository(cfg config.AppConfig, db *gorm.DB) repository.UserRepository {
	return &userRepository{
		cfg,
		db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, body *entity.User) (*entity.User, error) {
	hashedPassword, err := util.HashPassword(body.Password)
	if err != nil {
		log.Printf("[ERROR] User Repository - CreateUser : %v\n", err)
		return nil, err
	}

	user := entity.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Role:     entity.RoleUser,
	}

	if err := r.db.Model(&entity.User{}).
		Create(&user).
		Error; err != nil {
		log.Printf("[ERROR] User Repository - CreateUser : %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var result entity.User
	var total int64

	if err := r.db.Model(&entity.User{}).
		Where("email = ?", email).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] User Repository - GetUserByEmail : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("user with email %s does not exists", email)
		log.Printf("[ERROR] User Repository - GetUserByEmail : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var result entity.User
	var total int64

	if err := r.db.Model(&entity.User{}).
		Where("username = ?", username).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] User Repository - GetUserByEmail : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("user with username %s does not exists", username)
		log.Printf("[ERROR] User Repository - GetUserByEmail : %v\n", err)
		return nil, err
	}

	return &result, nil
}
