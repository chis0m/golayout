package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID  `json:"id,omitempty"`
	UserId       uint64     `json:"user_id,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	UserAgent    string     `json:"user_agent,omitempty"`
	ClientIP     string     `json:"client_ip,omitempty"`
	IsBlocked    bool       `json:"is_blocked,omitempty"`
	ExpiresAt    time.Time  `json:"expires_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
