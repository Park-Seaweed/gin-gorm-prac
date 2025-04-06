package dto

type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
