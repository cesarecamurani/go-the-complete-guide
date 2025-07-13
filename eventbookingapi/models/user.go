package models

import (
	"errors"
	"example.com/eventbookingapi/db"
	"example.com/eventbookingapi/utils"
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *User) Save() error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	query := `
	INSERT INTO users (id, username, email, password) 
	VALUES (?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = statement.Exec(u.ID, u.Username, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE username = ?`

	row := db.DB.QueryRow(query, u.Username)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("invalid credentials")
	}

	return nil
}
