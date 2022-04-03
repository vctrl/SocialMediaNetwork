package model

import (
	"context"
	"errors"
	"github.com/vctrl/SocialNetworkHighload/api/internal/db"
	"github.com/vctrl/SocialNetworkHighload/api/internal/password"
	"github.com/vctrl/SocialNetworkHighload/api/internal/session"
)

// afk for 15 min
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

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

type User struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Interests string `json:"interests"`
	City      string `json:"city"`
}

type Model interface {
	Login(ctx context.Context, r *LoginRequest) (user *User, token string, err error)
	Register(ctx context.Context, r *RegisterRequest) (id string, token string, err error)
	Logout(ctx context.Context) error
	LogoutAll(ctx context.Context) error

	GetUserInfo(ctx context.Context) error
	UpdateUserInfo(ctx context.Context) error
	DeleteUser(ctx context.Context) error

	GetFriends(ctx context.Context) error
	SendFriendRequest(ctx context.Context) error
	GetSentRequests(ctx context.Context) error
	GetIncomeRequests(ctx context.Context) error
	AcceptFriendRequest(ctx context.Context) error

	DeleteFriend(ctx context.Context) error
	CancelRequest(ctx context.Context) error
}

func New(users db.Users,
	profiles db.Profiles,
	friends db.Friends,
	friendRequests db.FriendRequests,
	sm session.SessionManager,
	ph password.PasswordHasher) (Model, error) {
	return &ModelImpl{
		Users: users,
		// todo
		Profiles:       profiles,
		Friends:        friends,
		FriendRequests: friendRequests,
		Sm:             sm,
		Ph:             ph,
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

func (m *ModelImpl) Logout(ctx context.Context) error {
	return errors.New("not implemented")
}

func (m *ModelImpl) LogoutAll(ctx context.Context) error {
	return errors.New("not implemented")
}

func (m *ModelImpl) GetUserInfo(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) UpdateUserInfo(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) DeleteUser(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) GetFriends(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) SendFriendRequest(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) GetSentRequests(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) GetIncomeRequests(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) AcceptFriendRequest(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) DeleteFriend(ctx context.Context) error {
	return nil
}

func (m *ModelImpl) CancelRequest(ctx context.Context) error {
	return nil
}
