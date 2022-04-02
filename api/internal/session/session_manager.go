package session

import (
	"context"
	"crypto/rsa"
	"github.com/SermoDigital/jose/crypto"
)

type Session struct {
	UserID uint32
	ID     string
}

type SessionManager interface {
	Create(ctx context.Context) error
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

func (sm *SessionManagerJWT) Create(ctx context.Context) error {
	return nil
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
