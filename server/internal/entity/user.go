package entity

import "github.com/google/uuid"

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username string
	Email    string
	Password string
	Role     Role `gorm:"type:role"`
}
