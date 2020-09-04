package repository

import (
	"errors"
	"github.com/SleepingNext/flick-backend-test/model"
)

func InsertSession(params *model.Session) (*model.Session, error) {
	query := `INSERT INTO sessions (user_id, device_id, status) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, params.UserID, params.DeviceID, params.Status).Scan(&params.ID)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func GetOneSession(id int64) (*model.Session, error) {
	var session model.Session

	query := `SELECT * FROM sessions WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&session.ID, session.UserID, session.DeviceID, session.Status)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func UpdateSession(params *model.Session) (*model.Session, error) {
	query := `UPDATE sessions SET user_id = $1, device_id = $2, status = $3 WHERE id = $4`
	res, err := db.Exec(query, params.UserID, params.DeviceID, params.Status, params.ID)
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

	return params, nil
}

func UpdateSessions(params *model.Session) error {
	query := `UPDATE sessions SET status = $1 WHERE id != $2 AND user_id = $3 AND status != $4`
	res, err := db.Exec(query, "INACTIVE", params.ID, params.UserID, "INACTIVE")
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count <= 0 {
		return errors.New("sql: no rows found")
	}

	return nil
}