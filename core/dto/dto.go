package dto

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=8,max=16" example:"12345678"`
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=8,max=16" example:"12345678"`
}

type SignUpResponse struct {
	Message string         `json:"message" example:"Sign Up successfully.!"`
	Token   *TokenResponse `json:"token" `
}
