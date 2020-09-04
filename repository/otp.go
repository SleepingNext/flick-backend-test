package repository

import (
	"github.com/SleepingNext/flick-backend-test/model"
)

func InsertOTP(params *model.OTP) (*model.OTP, error) {
	query := `INSERT INTO otps (user_id, code, ttl) VALUES ($1, $2, $3) RETURNING "id"`
	err := db.QueryRow(query, params.UserID, params.Code, params.Ttl).Scan(&params.ID)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func GetOneOTP(params *model.OTP) (*model.OTP, error) {
	var otp model.OTP

	query := `SELECT * FROM otps WHERE user_id = $1 AND code = $2`
	err := db.QueryRow(query, params.UserID, params.Code).Scan(&otp.ID, &otp.UserID, &otp.Code, &otp.Ttl)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}