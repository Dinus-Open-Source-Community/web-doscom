package dto

type RefreshTokenSuccessResponse struct {
	Message string `json:"message" example:"refresh token success"`
	Token   string `json:"token" example:"new_access_token_here"`
}

type RefreshTokenErrorResponse struct {
	Error   string `json:"error" example:"token expired"`
	Message string `json:"message" example:"refresh token invalid or expired"`
}
