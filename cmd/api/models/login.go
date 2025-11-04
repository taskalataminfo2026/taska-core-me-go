package models

type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Message      string `json:"message,omitempty"`
}

type LoginResponseParams struct {
	AccessToken  string
	RefreshToken string
	Message      string
}
