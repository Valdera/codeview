package util

import (
	"codeview/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	privateKey []byte
	publicKey  []byte
}

type JWTClaims struct {
	jwt.StandardClaims
	Id   string `json:"id"`
	Role string `json:"role"`
}

func NewJWTManager(cfg config.AppConfig) *JWTManager {
	return &JWTManager{
		privateKey: cfg.JWTPrivateKey,
		publicKey:  cfg.JWTPublicKey,
	}
}

func (manager *JWTManager) GenerateJwt(id string, tokenDuration time.Duration) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "user-issuer",
		},
		Id:   id,
		Role: "user",
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(manager.privateKey)
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
}

func (manager *JWTManager) Verify(accessToken string) (*JWTClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(manager.publicKey)
	if err != nil {
		return nil, fmt.Errorf("cannot parse key: %w", err)
	}

	token, err := jwt.ParseWithClaims(
		accessToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return key, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("cannot parse claims")
	}

	return claims, nil
}
