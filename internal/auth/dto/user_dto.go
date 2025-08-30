package dto

import "time"

type UserDTO struct {
	ID        string
	Email     string
	Username  string
	CreatedAt time.Time
}
