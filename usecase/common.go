package usecase

import "github.com/xlzd/gotp"

var otpGenerator *gotp.TOTP

func InitiateOTPGenerator() {
	otpGenerator = gotp.NewTOTP(secret, 6, OTP_TTL, nil)
}