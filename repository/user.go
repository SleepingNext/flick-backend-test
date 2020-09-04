package repository

import (
	"errors"
	"github.com/SleepingNext/flick-backend-test/model"
)

func InsertUser(params *model.User) (*model.User, error) {
	query := `INSERT INTO users (full_name, phone_number, email, password, status, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(query , params.FullName, params.PhoneNumber, params.Email, params.Password, params.Status, params.Role).Scan(&params.ID)
	if err != nil {
		return nil, err
	}

	params.Password = ""

	return params, nil
}

func GetOneUser(id int64) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.FullName, &user.PhoneNumber, &user.Email, &user.Password, &user.Status, &user.Role)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func GetOneUserByEmail(email string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.FullName, &user.PhoneNumber, &user.Email, &user.Password, &user.Status, &user.Role)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func GetOneUserByEmailAndPassword(params *model.User) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE email = $1 AND password = $2`
	err := db.QueryRow(query, params.Email, params.Password).Scan(&user.ID, &user.FullName, &user.PhoneNumber, &user.Email, &user.Password, &user.Status, &user.Role)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func UpdateUser(params *model.User) (*model.User, error) {
	query := `UPDATE users SET full_name = $1, phone_number = $2, email = $3, password = $4, status = $5, role = $6 WHERE id = $7`
	res, err := db.Exec(query, params.FullName, params.PhoneNumber, params.Email, params.Password, params.Status, params.Role, params.ID)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, errors.New("sql: no rows found")
	}

	params.Password = ""

	return params, nil
}
