package usecase

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SleepingNext/flick-backend-test/common"
	"github.com/SleepingNext/flick-backend-test/model"
	"github.com/SleepingNext/flick-backend-test/repository"
)

func RegisterUser(request *model.User) (*model.User, error) {
	user, err := repository.GetOneUserByEmail(request.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if user == nil {
		request.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(request.Password)))
		request.Status = "OTP NOT VERIFIED"
		request.Role = "COMMON"

		user, err = repository.InsertUser(request)
		if err != nil {
			return nil, err
		}
	} else if user.Status == "INACTIVE" {
		return user, errors.New("this email is already in use")
	}

	otp, err := GenerateOTP(user.ID)
	if err != nil {
		return nil, err
	}

	to := []string{user.Email}
	var cc []string
	subject := "OTP for Verifying Back-In User Registration"
	message := "Hello. This is your OTP: " + otp.Code + ".\nPlease use before: " + otp.Ttl.String()

	err = common.SendMail(to, cc, subject, message)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func RespondUserRegistration(request *model.User) (*model.User, error) {
	user, err := repository.GetOneUser(request.ID)
	if err != nil {
		return nil, err
	}

	if user.Status == "ACTIVE" || user.Status == "REJECTED"{
		return user, errors.New("this user registration is already responded")
	}

	if request.Status == "ACTIVE" {
		user.Status = request.Status

		user, err = repository.UpdateUser(user)
		if err != nil {
			return nil, err
		}

		to := []string{user.Email}
		var cc []string
		subject := "Your Registration Has Been Accepted!"
		message := "Hello. Congratulations on being accepted as the user of Bank-In!"

		err = common.SendMail(to, cc, subject, message)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
	user.Status = request.Status

	user, err = repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	to := []string{user.Email}
	var cc []string
	subject := "We are so sorry."
	message := "Hello. We are so sorry that we have to reject your registration due to false data given."

	err = common.SendMail(to, cc, subject, message)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(request *model.LoginRequest) (*model.User, error) {
	request.User.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(request.User.Password)))
	user, err := repository.GetOneUserByEmailAndPassword(request.User)
	if err != nil {
		return nil, err
	}

	session := &model.Session{
		UserID: request.User.ID,
		DeviceID: request.DeviceID,
		Status: "OTP NOT VERIFIED",
	}

	session, err = repository.InsertSession(session)
	if err != nil {
			return nil, err
		}

	otp, err := GenerateOTP(user.ID)
	if err != nil {
		return nil, err
	}

	to := []string{user.Email}
	var cc []string
	subject := "OTP for Verifying Back-In User Login"
	message := "Hello. This is your OTP: " + otp.Code + ".\nPlease use before: " + otp.Ttl.String()

	err = common.SendMail(to, cc, subject, message)
	if err != nil {
		return nil, err
	}

	return user, nil
}