package main

import "crypto/rsa"

type User struct {
	id           int
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken interface{}
}

type SecretKeys struct {
	Public  *rsa.PublicKey
	Private *rsa.PrivateKey
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Config struct {
	id               int
	AccessTokenTime  int64
	RefreshTokenTime int64
}
