package model

import (
	"context"
	"time"
)

func (m *ModelImpl) Login(ctx context.Context, r *LoginRequest) (user *User, token string, err error) {
	dbUser, err := m.Users.GetByLogin(ctx, r.Login)
	if err != nil {
		return nil, "", err
	}

	dbProfile, err := m.Profiles.GetByUserID(ctx, dbUser.ID)
	if err != nil {
		return nil, "", err
	}

	// todo move to config
	expiresAt := time.Now().Add(time.Hour * 5)
	token, err = m.Sm.Create(ctx, dbUser.ID, dbUser.Login, expiresAt.UnixMicro())
	if err != nil {
		return nil, "", err
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
