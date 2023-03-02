package entity

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
	Role     Role `gorm:"type:role"`
}
