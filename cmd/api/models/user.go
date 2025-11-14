package models

import (
	"encoding/json"
	"errors"
	"regexp"
	"taska-core-me-go/cmd/api/constants"
	"time"
)

// Objecto principal.
type User struct {
	ID                   int64
	UserName             string
	FirstName            string
	LastName             string
	Email                string
	CountryCode          string
	PhoneNumber          string
	PasswordHash         string
	UserType             string
	ProfilePictureURL    string
	Bio                  string
	BirthDate            time.Time
	Gender               string
	Specialties          json.RawMessage
	AccountType          string
	TasksCompleted       int64
	ThemePreference      string
	VerificationCode     string
	VerificationAttempts int64
	CodeExpiration       string
	MessageID            string
	IsActive             bool
	IsVerified           bool
	IsBlocked            bool
	LastLoginAt          time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type ParamUser struct {
	ID                   int32           `json:"id"`
	UserName             string          `json:"user_name"`
	StatusDetails        string          `json:"status_details"`
	StatusCode           int64           `json:"status_code"`
	FirstName            string          `json:"firstName"`
	LastName             string          `json:"lastName"`
	Email                string          `json:"email"`
	PhoneNumber          string          `json:"phoneNumber"`
	PasswordHash         string          `json:"password_hash"`
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

type CreateUserRequest struct {
	UserName        string
	FirstName       string
	LastName        string
	NumberExtension string
	PhoneNumber     string
	Email           string
	PasswordHash    string
	ConfirmPassword string
	Gender          string
	UserType        string
	Roles           []string
	Skills          map[string]interface{}
}

func (u *CreateUserRequest) IsValidatePasswordStrength() error {
	if len(u.PasswordHash) < 5 {
		return errors.New(constants.MsgPasswordMinLength)
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString

	if !hasUpper(u.PasswordHash) {
		return errors.New(constants.MsgPasswordUppercase)
	}

	if !hasLower(u.PasswordHash) {
		return errors.New(constants.MsgPasswordLowercase)
	}

	if !hasNumber(u.PasswordHash) {
		return errors.New(constants.MsgPasswordNumber)
	}

	if !hasSpecial(u.PasswordHash) {
		return errors.New(constants.MsgPasswordSpecialChar)
	}

	return nil
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u *CreateUserRequest) IsValidEmail() error {
	if !emailRegex.MatchString(u.Email) {
		return errors.New(constants.MsgInvalidEmail)
	}
	return nil
}

func (u *CreateUserRequest) ValidateUsername() error {
	if len(u.UserName) < 3 {
		return errors.New(constants.MsgUsernameMinLength)
	}

	if len(u.UserName) > 75 {
		return errors.New(constants.MsgUsernameMaxLength)
	}

	// Solo permite letras, números, guiones bajos y puntos
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_.]+$`).MatchString
	if !validUsername(u.UserName) {
		return errors.New(constants.MsgUsernameInvalidChars)
	}

	return nil
}

type BodyUserResponse struct {
	AccessToken  string
	RefreshToken string
	Data         UserResponse
}

type UserResponse struct {
	ID                   int64
	UserName             string
	FirstName            string
	LastName             string
	Email                string
	CountryCode          string
	PhoneNumber          string
	PasswordHash         string // no se expone
	UserType             string
	ProfilePictureURL    string
	Bio                  string
	BirthDate            time.Time
	Gender               string
	Specialties          json.RawMessage
	AccountType          string
	TasksCompleted       int64
	ThemePreference      string
	VerificationCode     string
	VerificationAttempts int64
	IsActive             bool
	IsVerified           bool
	IsBlocked            bool
	LastLoginAt          time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (u *User) Load() UserResponse {
	return UserResponse{
		ID:                   u.ID,
		UserName:             u.UserName,
		FirstName:            u.FirstName,
		LastName:             u.LastName,
		Email:                u.Email,
		CountryCode:          u.CountryCode,
		PhoneNumber:          u.PhoneNumber,
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
		IsActive:             u.IsActive,
		IsVerified:           u.IsVerified,
		IsBlocked:            u.IsBlocked,
		LastLoginAt:          u.LastLoginAt,
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
	}
}

type UpdateUserRequest struct {
	ID                int64
	UserName          string
	FirstName         string
	LastName          string
	Bio               string
	Gender            string
	ThemePreference   string
	ProfilePictureURL string
	BirthDate         string
	Specialties       json.RawMessage
}
