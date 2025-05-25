package models

import (
	"errors"

	"example.com/go-basic-backend/db"
	"example.com/go-basic-backend/utils"
)

type User struct {
	UserID   int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
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

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users where email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retreivedPassword string
	err := row.Scan(&u.UserID, &retreivedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isCorrectPassword := utils.CompareHashPassword(u.Password, retreivedPassword)
	if !isCorrectPassword {
		return errors.New("invalid credentials")
	}

	return nil
}
