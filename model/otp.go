package model

import "time"

type OTP struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"user_id"`
	Code   string     `json:"code"`
	Ttl    *time.Time `json:"ttl"`
}
