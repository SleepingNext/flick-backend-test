package delivery

import (
	"encoding/json"
	"github.com/SleepingNext/flick-backend-test/model"
	"github.com/SleepingNext/flick-backend-test/usecase"
	"io/ioutil"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		byteData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var tempUser model.User
		err = json.Unmarshal(byteData, &tempUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := usecase.RegisterUser(&tempUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func VerifyUserRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		byteData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var tempOTP model.OTP
		err = json.Unmarshal(byteData, &tempOTP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := usecase.VerifyOTPForUserRegistration(&tempOTP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func RespondUserRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		byteData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var tempUser model.User
		err = json.Unmarshal(byteData, &tempUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := usecase.RespondUserRegistration(&tempUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		byteData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var request model.LoginRequest
		err = json.Unmarshal(byteData, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := usecase.Login(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func VerifyLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		byteData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var tempOTP model.OTP
		err = json.Unmarshal(byteData, &tempOTP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := usecase.VerifyOTPForLogin(&tempOTP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}