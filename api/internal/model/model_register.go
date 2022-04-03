package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/vctrl/SocialNetworkHighload/api/internal/db"
	"time"
)

func (m *ModelImpl) Register(ctx context.Context, r *RegisterRequest) (userID, token string, err error) {
	id := uuid.New()
	passwordHash := m.Ph.HashPass(r.Password)

	dbUser := db.NewUser(id.String(), r.Login, passwordHash, time.Now())
	err = m.Users.Add(ctx, dbUser)

	if err != nil {
		return "", "", err
	}

	// todo move to config
	expiresAt := time.Now().Add(time.Hour * 5)
	token, err = m.Sm.Create(ctx, id.String(), r.Login, expiresAt.UnixMicro())
	if err != nil {
		return "", "", err
	}

	return id.String(), token, nil
}
