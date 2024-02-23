package types

import (
	"time"
)

type UserResponse struct {
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
