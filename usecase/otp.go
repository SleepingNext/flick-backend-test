package usecase

import (
	"database/sql"
	"errors"
	"github.com/SleepingNext/flick-backend-test/model"
	"github.com/SleepingNext/flick-backend-test/repository"
	"time"
)

const OTP_TTL = 60
const secret = "4S62BZNFXXSZLCRO"

func GenerateOTP(userID int64) (*model.OTP, error) {
	ttl := time.Now().Local().Add(time.Second * OTP_TTL)
	otp := &model.OTP{
		Code:   otpGenerator.At(int(ttl.Unix())),
		UserID: userID,
		Ttl:    &ttl,
	}

	otp, err := repository.InsertOTP(otp)
	if err != nil {
		return nil, err
	}

	return otp, nil
}

func VerifyOTPForUserRegistration(request *model.OTP) (*model.User, error) {
	user, err := repository.GetOneUser(request.UserID)
	if err != nil {
		return nil, err
	}

	if user.Status == "INACTIVE" {
		return user, errors.New("this user is already verified")
	}

	notValidError := errors.New("OTP is not valid")

	otp, err := repository.GetOneOTP(request)
	if err != nil {
		return nil, err
	}

	if request.Code != otp.Code {
		return nil, notValidError
	}

	isValid := otpGenerator.Verify(otp.Code, int(otp.Ttl.Unix()))
	if isValid {
		user.Status = "INACTIVE"

		user, err = repository.UpdateUser(user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, notValidError
}

func VerifyOTPForLogin(request *model.OTP) (*model.Session, error) {
	session, err := repository.GetOneSession(request.UserID)
	if err != nil {
		return nil, err
	}

	if session.Status == "ACTIVE" {
		return session, errors.New("this login is already verified")
	}

	notValidError := errors.New("OTP is not valid")

	otp, err := repository.GetOneOTP(request)
	if err != nil {
		return nil, err
	}

	if request.Code != otp.Code {
		return nil, notValidError
	}

	isValid := otpGenerator.Verify(otp.Code, int(otp.Ttl.Unix()))
	if isValid {
		session.Status = "ACTIVE"

		err = repository.UpdateSessions(session)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		session, err = repository.UpdateSession(session)
		if err != nil {
			return nil, err
		}

		return session, nil
	}

	return nil, notValidError
}
