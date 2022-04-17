package session

import (
	"context"
	"crypto/rsa"
	"github.com/SermoDigital/jose/crypto"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/vctrl/social-media-network/api/internal/config"
	"github.com/xhit/go-str2duration/v2"
	"io/ioutil"
	"log"
	"time"
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
	accessTTL  time.Duration
}

func FromConfig(cfg *config.Config) (SessionManager, error) {
	publicKey, err := ioutil.ReadFile(cfg.Token.PublicKeyPath)
	if err != nil {
		log.Fatalf("failed to read public key: %v", err)
	}

	privateKey, err := ioutil.ReadFile(cfg.Token.PrivateKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	ttl, err := str2duration.ParseDuration(cfg.Token.AccessTTL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse access ttl")
	}

	return NewSessionsJWTManager(privateKey, publicKey, ttl)
}

func NewSessionsJWTManager(privateKeyBytes, publicKeyBytes []byte, accessTTL time.Duration) (SessionManager, error) {
	privateKey, err := crypto.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse rsa private key")
	}
	publicKey, err := crypto.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse rsa public key")
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
		return "", errors.Wrap(err, "sign token with private key")
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
