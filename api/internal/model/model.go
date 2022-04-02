package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/vctrl/SocialNetworkHighload/api/internal/db"
	"github.com/vctrl/SocialNetworkHighload/api/internal/password"
	"github.com/vctrl/SocialNetworkHighload/api/internal/session"
	"time"
)

type RegisterRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Interests string `json:"interests"`
	City      string `json:"city"`
}

type Model interface {
	Register(ctx context.Context, r *RegisterRequest) error

	GetUser(ctx context.Context) error
	UpdateUser(ctx context.Context) error
	DeleteUser(ctx context.Context) error

	CreateSession(ctx context.Context) error
	CheckSession(ctx context.Context) error
	DeleteSession(ctx context.Context) error
	DeleteAll(ctx context.Context) error

	// TODO friends and friend requests
	//AddFriend(ctx context.Context)
}

func New(users db.Users, privateKeyBytes, publicKeyBytes []byte) (Model, error) {
	sm, err := session.NewSessionsJWTManager(privateKeyBytes, publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return &ModelImpl{
		Users: users,
		// todo
		Profiles:       nil,
		Friends:        nil,
		FriendRequests: nil,
		Sm:             sm,
		Ph:             nil,
	}, nil
}

type ModelImpl struct {
	Users          db.Users
	Profiles       db.Profiles
	Friends        db.Friends
	FriendRequests db.FriendRequests

	Sm session.SessionManager

	Ph password.PasswordHasher
}

func (m *ModelImpl) Register(ctx context.Context, r *RegisterRequest) error {
	id := uuid.New()
	passwordHash, err := m.Ph.HashPass(r.Password)

	if err != nil {
		return err
	}

	dbUser := db.NewUser(id.String(), r.Login, passwordHash, time.Now())
	err = m.Users.Add(ctx, dbUser)

	if err != nil {
		return err
	}

	return nil
}

func (m *ModelImpl) GetUser(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) UpdateUser(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) DeleteUser(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) CreateSession(ctx context.Context) error {
	return nil
}
func (m *ModelImpl) CheckSession(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) DeleteSession(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) DeleteAll(ctx context.Context) error {
	return nil
}
