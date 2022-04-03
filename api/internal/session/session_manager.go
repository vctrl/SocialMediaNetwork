package session

import (
	"context"
	"crypto/rsa"
	"github.com/SermoDigital/jose/crypto"
	"github.com/dgrijalva/jwt-go"
)

type Session struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
	jwt.StandardClaims
}

type SessionManager interface {
	Create(ctx context.Context, userID, login string, expiresAt int64) (string, error)
	Check(ctx context.Context) error
	Destroy(ctx context.Context) error
	DestroyAll(ctx context.Context) error
}

type SessionManagerJWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewSessionsJWTManager(privateKeyBytes, publicKeyBytes []byte) (SessionManager, error) {
	privateKey, err := crypto.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	publicKey, err := crypto.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return &SessionManagerJWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (sm *SessionManagerJWT) Create(ctx context.Context, userID, login string, expiresAt int64) (string, error) {
	sess := &Session{
		UserID: userID,
		Login:  login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, sess)
	signed, err := token.SignedString(sm.privateKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (sm *SessionManagerJWT) Check(ctx context.Context) error {
	return nil
}

func (sm *SessionManagerJWT) Destroy(ctx context.Context) error {
	// ¯\_(ツ)_/¯
	return nil
}

func (sm *SessionManagerJWT) DestroyAll(ctx context.Context) error {
	// ¯\_(ツ)_/¯
	return nil
}
