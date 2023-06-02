package entity

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
	Role     Role `gorm:"type:role"`
}
