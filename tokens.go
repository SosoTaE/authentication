package main

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func generateToken(p *User, secretKey *rsa.PrivateKey, tokenTime int64) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "authMicroService",
		Subject:   p.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenTime) * time.Second)), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(secretKey)

	return signedToken, err

}

func generateTokens(p *User, tokenSecretKey *SecretKeys, tokensTime *Config) (Tokens, error) {
	accessToken, err := generateToken(p, tokenSecretKey.Private, tokensTime.AccessTokenTime)
	if err != nil {
		return Tokens{}, err
	}

	refreshToken, err := generateToken(p, tokenSecretKey.Private, tokensTime.RefreshTokenTime)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
