package models

import (
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (BlacklistedTokenDb) TableName() string { return "blacklisted_tokens" }

type BlacklistedTokenDb struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	UserID    int64     `gorm:"column:user_id"`
	Token     string    `gorm:"column:token"`
	TokenType string    `gorm:"column:token_type"`
	Reason    string    `gorm:"column:reason"`
	IPAddress string    `gorm:"column:ip_address"`
	UserAgent string    `gorm:"column:user_agent"`
	RevokedAt time.Time `gorm:"column:revoked_at"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (db *BlacklistedTokenDb) ToDomainModel() models.BlacklistedToken {
	return models.BlacklistedToken{
		ID:        db.ID,
		UserID:    db.UserID,
		Token:     db.Token,
		TokenType: db.TokenType,
		Reason:    db.Reason,
		IPAddress: db.IPAddress,
		UserAgent: db.UserAgent,
		RevokedAt: db.RevokedAt,
		ExpiresAt: db.ExpiresAt,
		CreatedAt: db.CreatedAt,
	}
}

func (db *BlacklistedTokenDb) Load(m models.BlacklistedToken) {
	db.ID = m.ID
	db.UserID = m.UserID
	db.Token = m.Token
	db.TokenType = m.TokenType
	db.Reason = m.Reason
	db.IPAddress = m.IPAddress
	db.UserAgent = m.UserAgent
	db.RevokedAt = m.RevokedAt
	db.ExpiresAt = m.ExpiresAt
	db.CreatedAt = m.CreatedAt
}
