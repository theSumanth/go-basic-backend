package models

import (
	"example.com/go-basic-backend/db"
	"example.com/go-basic-backend/utils"
)

type User struct {
	UserID   int64
	Email    string
	Password string
}

func (u User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
