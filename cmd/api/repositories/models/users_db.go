package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (UserDb) TableName() string { return "users" }

type UserDb struct {
	ID                   int64           `gorm:"column:id"`
	UserName             string          `gorm:"column:user_name"`
	FirstName            string          `gorm:"column:first_name"`
	LastName             string          `gorm:"column:last_name"`
	Email                string          `gorm:"column:email"`
	CountryCode          string          `gorm:"column:country_code"`
	PhoneNumber          string          `gorm:"column:phone_number"`
	PasswordHash         string          `gorm:"column:password_hash"`
	UserType             string          `gorm:"column:user_type"`
	ProfilePictureURL    string          `gorm:"column:profile_picture_url"`
	Bio                  string          `gorm:"column:bio"`
	BirthDate            time.Time       `gorm:"column:birth_date"`
	Gender               string          `gorm:"column:gender"`
	Specialties          json.RawMessage `gorm:"column:specialties"`
	AccountType          string          `gorm:"column:account_type"`
	TasksCompleted       int64           `gorm:"column:tasks_completed"`
	ThemePreference      string          `gorm:"column:theme_preference"`
	VerificationCode     string          `gorm:"column:verification_code"`
	VerificationAttempts int64           `gorm:"column:verification_attempts"`
	CodeExpiration       string          `gorm:"column:code_expiration"`
	MessageID            string          `gorm:"column:message_id"`
	IsActive             bool            `gorm:"column:is_active"`
	IsVerified           bool            `gorm:"column:is_verified"`
	IsBlocked            bool            `gorm:"column:is_blocked"`
	LastLoginAt          time.Time       `gorm:"type:timestamp" json:"last_login_at,omitempty"`
	CreatedAt            time.Time       `gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt            time.Time       `gorm:"type:timestamp" json:"updated_at,omitempty"`
}

func (u *UserDb) ToDomainModel() models.User {
	return models.User{
		ID:                   u.ID,
		UserName:             u.UserName,
		FirstName:            u.FirstName,
		LastName:             u.LastName,
		Email:                u.Email,
		CountryCode:          u.CountryCode,
		PhoneNumber:          u.PhoneNumber,
		PasswordHash:         u.PasswordHash,
		UserType:             u.UserType,
		ProfilePictureURL:    u.ProfilePictureURL,
		Bio:                  u.Bio,
		BirthDate:            u.BirthDate,
		Gender:               u.Gender,
		Specialties:          u.Specialties,
		AccountType:          u.AccountType,
		TasksCompleted:       u.TasksCompleted,
		ThemePreference:      u.ThemePreference,
		VerificationCode:     u.VerificationCode,
		VerificationAttempts: u.VerificationAttempts,
		CodeExpiration:       u.CodeExpiration,
		MessageID:            u.MessageID,
		IsActive:             u.IsActive,
		IsVerified:           u.IsVerified,
		IsBlocked:            u.IsBlocked,
		LastLoginAt:          u.LastLoginAt,
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
	}
}

func (u *UserDb) Load(item models.User) {
	u.ID = item.ID
	u.UserName = item.UserName
	u.FirstName = item.FirstName
	u.LastName = item.LastName
	u.Email = item.Email
	u.CountryCode = item.CountryCode
	u.PhoneNumber = item.PhoneNumber
	u.PasswordHash = item.PasswordHash
	u.UserType = item.UserType
	u.ProfilePictureURL = item.ProfilePictureURL
	u.Bio = item.Bio
	u.BirthDate = item.BirthDate
	u.Gender = item.Gender
	u.Specialties = item.Specialties
	u.AccountType = item.AccountType
	u.TasksCompleted = item.TasksCompleted
	u.ThemePreference = item.ThemePreference
	u.VerificationCode = item.VerificationCode
	u.VerificationAttempts = item.VerificationAttempts
	u.IsActive = item.IsActive
	u.IsVerified = item.IsVerified
	u.IsBlocked = item.IsBlocked
	u.LastLoginAt = item.LastLoginAt
	u.CreatedAt = item.CreatedAt
	u.UpdatedAt = item.UpdatedAt
}

type ParamUserDB struct {
	ID                   int32           `json:"id"`
	UserName             string          `json:"user_name""`
	FirstName            string          `json:"firstName"`
	LastName             string          `json:"lastName"`
	Email                string          `json:"email"`
	PhoneNumber          string          `json:"phoneNumber"`
	UserType             string          `json:"userType"`
	ProfilePictureURL    string          `json:"profilePictureUrl"`
	Bio                  string          `json:"bio,omitempty"`
	Gender               string          `json:"gender,omitempty"`
	Specialties          json.RawMessage `json:"specialties,omitempty"`
	AccountType          string          `json:"accountType,omitempty"`
	TasksCompleted       int32           `json:"tasksCompleted"`
	ThemePreference      string          `json:"themePreference"`
	VerificationCode     string          `json:"verificationCode,omitempty"`
	VerificationAttempts int32           `json:"verificationAttempts"`
	IsActive             bool            `json:"isActive,omitempty"`
	IsVerified           bool            `json:"isVerified,omitempty"`
	IsBlocked            bool            `json:"isBlocked,omitempty"`
}

func (u *ParamUserDB) GetQueryUsers() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if u.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, u.ID)
	}

	if u.UserName > "" {
		query = append(query, "user_name = ? ")
		params = append(params, u.UserName)
	}

	if u.Email > "" {
		query = append(query, "email = ? ")
		params = append(params, u.Email)
	}

	if u.Gender != "" {
		query = append(query, "gender = ? ")
		params = append(params, u.Gender)
	}

	return strings.Join(query, " AND "), params
}

func (db *ParamUserDB) ToDB(u *models.ParamUser) {
	db.ID = u.ID
	db.UserName = u.UserName
	db.FirstName = u.FirstName
	db.LastName = u.LastName
	db.Email = u.Email
	db.PhoneNumber = u.PhoneNumber
	db.UserType = u.UserType
	db.ProfilePictureURL = u.ProfilePictureURL
	db.Bio = u.Bio
	db.Gender = u.Gender
	db.Specialties = u.Specialties
	db.AccountType = u.AccountType
	db.TasksCompleted = u.TasksCompleted
	db.ThemePreference = u.ThemePreference
	db.VerificationCode = u.VerificationCode
	db.VerificationAttempts = u.VerificationAttempts
	db.IsActive = u.IsActive
	db.IsVerified = u.IsVerified
	db.IsBlocked = u.IsBlocked
}

func (u *UserDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (u *UserDb) BeforeUpdate(tx *gorm.DB) (err error) {
	u.LastLoginAt = time.Now().Local()
	return
}
