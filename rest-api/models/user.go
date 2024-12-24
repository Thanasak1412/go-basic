package models

import (
	"database/sql"
	"github.com/Thanasak1412/go-rest-api/db"
	"github.com/Thanasak1412/go-rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

func (u *User) Save() error {
	stmt, err := db.Db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashedPassword

	_, err = stmt.Exec(u.Email, u.Password)

	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	row := db.Db.QueryRow(query, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user User
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
