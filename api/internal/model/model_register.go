package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/vctrl/social-media-network/api/internal/db/mysql"
	"time"
)

func (m *ModelImpl) Register(ctx context.Context, r *RegisterRequest) (userID, token string, err error) {
	id := uuid.New()
	passwordHash := m.Ph.HashPass(r.Password)

	dbUser := mysql.NewUser(id.String(), r.Login, passwordHash, time.Now())
	err = m.Users.Add(ctx, dbUser)

	if err != nil {
		return "", "", errors.WithMessage(err, "create user")
	}

	// todo move to config
	expiresAt := time.Now().Add(time.Hour * 5)
	token, err = m.Sm.Create(ctx, id.String(), r.Login, expiresAt.UnixMicro())
	if err != nil {
		return "", "", errors.WithMessage(err, "create session")
	}

	return id.String(), token, nil
}
