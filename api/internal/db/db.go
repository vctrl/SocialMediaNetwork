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
	GetByLogin(ctx context.Context, login string) (*User, error)
}

type Profiles interface {
	GetByUserID(ctx context.Context, id string) (*Profile, error)
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
	query := "INSERT INTO users (`id`, `login`, `password`, `created_at`) VALUES (?, ?, ?, ?)"
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

func (r *UsersMySQL) GetByLogin(ctx context.Context, login string) (*User, error) {
	query := "SELECT * FROM users WHERE login = ?"

	res := r.db.QueryRow(query, login)

	u := &User{}
	err := res.Scan(&u.ID, &u.Login, &u.Password, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return u, nil
}

type ProfilesMySQL struct {
	db *sql.DB
}

func NewProfilesMySQL(db *sql.DB) Profiles {
	return &ProfilesMySQL{db: db}
}

func (r *ProfilesMySQL) GetByUserID(ctx context.Context, id string) (*Profile, error) {
	query := "SELECT * FROM profiles WHERE user_id  = ?"

	res := r.db.QueryRow(query, id)
	p := &Profile{}
	err := res.Scan(&p.ID, &p.UserID, &p.Name, &p.Surname, &p.Sex, &p.Interests, &p.City)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return p, nil
}

type FriendsMySQL struct {
	db *sql.DB
}

func NewFriendsMySQL(db *sql.DB) Friends {
	return &FriendsMySQL{
		db: db,
	}
}

type FriendRequestsMySQL struct {
	db *sql.DB
}

func NewFriendRequestsMySQL(db *sql.DB) FriendRequests {
	return &FriendRequestsMySQL{
		db: db,
	}
}
