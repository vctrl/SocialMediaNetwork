package model

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func (m *ModelImpl) Login(ctx context.Context, r *LoginRequest) (user *User, token string, err error) {
	dbUser, err := m.Users.GetByLogin(ctx, r.Login)
	if err != nil {
		return nil, "", errors.WithMessage(err, "get user by login")
	}

	dbProfile, err := m.Profiles.GetByUserID(ctx, dbUser.ID)
	if err != nil {
		return nil, "", errors.WithMessage(err, "get profile by user id")
	}

	// todo move to config
	expiresAt := time.Now().Add(time.Hour * 5)
	token, err = m.Sm.Create(ctx, dbUser.ID, dbUser.Login, expiresAt.UnixMicro())
	if err != nil {
		return nil, "", errors.WithMessage(err, "create session")
	}

	return &User{
		ID:        dbUser.ID,
		Login:     dbUser.Login,
		Name:      dbProfile.Name,
		Surname:   dbProfile.Surname,
		Age:       dbProfile.Age,
		Sex:       dbProfile.Sex,
		Interests: dbProfile.Interests,
		City:      dbProfile.City,
	}, token, nil
}
