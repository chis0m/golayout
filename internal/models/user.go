package models

import "time"

type User struct {
	ID              uint64     `json:"id"`
	FirstName       *string    `json:"first_name,omitempty"`
	LastName        *string    `json:"last_name,omitempty"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `json:"emailVerified_at,omitempty"`
	PasswordHash    string     `json:"password_hash"`
	Address         *string    `json:"address,omitempty"`
	BVN             *string    `json:"bvn,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
