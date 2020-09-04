package model

type LoginRequest struct {
	User *User `json:"user" binding:"required"`
	DeviceID string `json:"device_id" binding:"required"`
}
