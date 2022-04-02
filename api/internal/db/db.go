package db

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Profile struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Interests string `json:"interests"`
	City      string `json:"city"`
	// todo timestamp
}

func NewUser(id, login string, password []byte, createdAt time.Time) *User {
	return &User{
		ID:        id,
		Login:     login,
		Password:  password,
		CreatedAt: createdAt,
	}
}

func NewProfile(id, userID, name, surname string, age int, sex, interests, city string) *Profile {
	return &Profile{
		ID:        id,
		UserID:    userID,
		Name:      name,
		Surname:   surname,
		Age:       age,
		Sex:       sex,
		Interests: interests,
		City:      city,
	}
}

type Users interface {
	Add(ctx context.Context, user *User) error
}

type Profiles interface {
}

type Friends interface {
}

type FriendRequests interface {
}

type UsersMySQL struct {
	db *sql.DB
}

func NewUsersMySQL(db *sql.DB) Users {
	return &UsersMySQL{
		db: db,
	}
}

func (r *UsersMySQL) Add(ctx context.Context, user *User) error {
	query := "INSERT INTO users (`id`, `login`, `password`, `created_at`) VALUES (?, ?)"
	res, err := r.db.Exec(query, user.ID, user.Login, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
