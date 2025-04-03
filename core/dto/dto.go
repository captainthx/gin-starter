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

type UserInfo struct {
	Id      uint   `json:"id" example:"2"`
	Email   string `json:"email" example:"user@example.com"`
	Avartar string `json:"avatar" example:"http://localhost:8080/v1/file/serve/fileName"`
}

type UpdateAvartarRequest struct {
	AvatarUrl string `json:"avatarUrl" validate:"required,url" example:"http://localhost:8080/v1/file/serve/filename"`
}

type UploadFileResponse struct {
	FileName string `json:"filename" example:"asdasdadsd2.png" `
	FileUrl  string `json:"fileUrl" example:"http://localhost:8080/v1/file/serve/fileName"`
	Size     int64  `json:"size" expple:"1000"`
}
