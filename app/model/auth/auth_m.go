package auth

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
