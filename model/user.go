package model

// User status will be like this: OTP NOT VERIFIED -> INACTIVE -> ACTIVE -> TEMPORARILY BLOCKED || OTP NOT VERIFIED -> INACTIVE -> REJECTED
type User struct {
	ID          int64  `json:"id"`
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"` // SHA256
	Status      string `json:"status"`
	Role        string `json:"role"`
}
