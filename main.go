package main

import (
	"database/sql"
	"fmt"
	"github.com/SleepingNext/flick-backend-test/delivery"
	"github.com/SleepingNext/flick-backend-test/repository"
	"github.com/SleepingNext/flick-backend-test/usecase"
	"log"
	"net/http"
)

var basePort = ":8080"

func OpenPostgresConnection() (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "sleepingnext", "kevin99123", "flick_backend_test"),
	)
}

func main() {
	db, err := OpenPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository.InitiateDB(db)
	usecase.InitiateOTPGenerator()

	http.HandleFunc("/user/register-user", delivery.RegisterUser)
	http.HandleFunc("/user/verify-user-registration", delivery.VerifyUserRegistration)
	http.HandleFunc("/user/respond-user-registration", delivery.RespondUserRegistration)
	http.HandleFunc("/user/login", delivery.Login)
	http.HandleFunc("/user/verify-login", delivery.VerifyLogin)

	fmt.Println("starting Back-In at http://localhost:8080/")
	http.ListenAndServe(basePort, nil)
}