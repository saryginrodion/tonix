package model

import (
	"github.com/jmoiron/sqlx"
)

type UserFeatures struct {
	DB *sqlx.DB
}

func Users(db *sqlx.DB) *UserFeatures {
	return &UserFeatures{DB: db}
}

func (f *UserFeatures) New(email string, password string, username string) (*string, error) {
	row := f.DB.QueryRow(
		`INSERT INTO users (email, password, username) VALUES ($1, $2, $3) RETURNING id`,
		email, password, username,
	)

	userId := ""

	err := row.Scan(&userId)

	if err != nil {
		return nil, err
	}

	return &userId, nil
}

func (f *UserFeatures) Update(user *User) error {
	_, err := f.DB.NamedExec(
		`UPDATE users
		SET email = :email
			password = :password
			username = :username
			displayable_name = :displayable_name
			description = :description
			email_verified = :email_verified
			balance = :balance
			withdrawal_balance = :withdrawal_balance
			last_withdrawal_card = :last_withdrawal_card
			avatar_id = :avatar_id
			identity_photo_id = :identity_photo_id
			WHERE id = :id`,
		user,
	)

	return err
}

func (f *UserFeatures) IsRegistered(username string, email string) (bool, error) {
	user := []User{}
	err := f.DB.Select(&user, "SELECT * FROM users WHERE username = $1 OR email = $2", username, email)
	return len(user) != 0, err
}

func (f *UserFeatures) ByUsername(username string) (*User, error) {
	user := User{}
	err := f.DB.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (f *UserFeatures) ByEmail(email string) (*User, error) {
	user := User{}
	err := f.DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (f *UserFeatures) ById(id string) (*User, error) {
	user := User{}
	err := f.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
