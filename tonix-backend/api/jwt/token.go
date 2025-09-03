package jwt

import (
	"time"
	"tonix/backend/api/utils"
)

type TokenType int

const (
	Access = iota
	Refresh
)

type TokenHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type TokenPayload[T any] struct {
	TokenType TokenType `json:"token_type"`
	TokenId   string    `json:"token_id"`
	Exp       int64     `json:"exp"`
	Data      T         `json:"data"`
}

type Token[T any] struct {
	Header  *TokenHeader
	Payload *TokenPayload[T]
}

func NewTokenPayload[T any](body T, expirationTime time.Time, tokenType TokenType, tokenId string) TokenPayload[T] {
	newToken := TokenPayload[T]{
		TokenType: tokenType,
		TokenId:   tokenId,
		Exp:       expirationTime.Unix(),
		Data:      body,
	}

	return newToken
}

func GenerateToken[T any](payload TokenPayload[T], secret string) (string, error) {
	header := TokenHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerJson, err := utils.ToJsonString(header)
	if err != nil {
		return "", err
	}

	headerB64 := utils.ToBase64(headerJson)

	tokenJson, err := utils.ToJsonString(payload)
	if err != nil {
		return "", err
	}

	tokenB64 := utils.ToBase64(tokenJson)
	tokenString := headerB64 + "." + tokenB64

	sign := utils.ToBase64(utils.HmacSha256(tokenString, secret))
	tokenString = tokenString + "." + sign

	return tokenString, nil
}
