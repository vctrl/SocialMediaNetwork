package mysql

import (
	"context"
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
