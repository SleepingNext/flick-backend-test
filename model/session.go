package model

type Session struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	DeviceID string `json:"device_id"`
	Status   string `json:"status"`
}
