package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int       `gorm:"AUTO_INCREMENT" json:"id"`
	Username  string    `form:"username" json:"username"`
	Email     string    `form:"email" json:"email"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
