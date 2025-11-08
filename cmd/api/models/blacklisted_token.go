package models

import "time"

type BlacklistedToken struct {
	ID        int64
	UserID    int64
	Token     string
	TokenType string
	Reason    string
	IPAddress string
	UserAgent string
	RevokedAt time.Time
	ExpiresAt time.Time
	CreatedAt time.Time
}
