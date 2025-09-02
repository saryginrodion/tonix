package jwt

import (
	"strings"
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

func ParseToken[T any](token string, secret string) (*Token[T], error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return nil, NewJWTError("JWT must contain 3 parts")
	}

	headerAndPayloadParts := parts[0] + "." + parts[1]
	sign := utils.ToBase64(utils.HmacSha256(headerAndPayloadParts, secret))

	if sign != parts[2] {
		return nil, NewJWTError("JWT sign check is failed")
	}

	headerJson, err := utils.FromBase64(parts[0])
	if err != nil {
		return nil, err
	}

	header, err := utils.FromJsonString[TokenHeader](headerJson)
	if err != nil {
		return nil, err
	}

	payloadJson, err := utils.FromBase64(parts[1])
	if err != nil {
		return nil, err
	}

	payload, err := utils.FromJsonString[TokenPayload[T]](payloadJson)
	if err != nil {
		return nil, err
	}

	return &Token[T]{
		Header:  header,
		Payload: payload,
	}, nil
}

func VerifyToken[T any](token *Token[T], tokenType TokenType) error {
	if token.Payload.Exp > time.Now().Unix() {
		return NewJWTError("JWT expired")
	}

	if token.Payload.TokenType != tokenType {
		return NewJWTError("JWT incorrect token type")
	}

	// TODO: Add check for blacklist with token.Payload.TokenId
	return nil
}

func ParseAndVerifyToken[T any](token string, tokenType TokenType, secret string) (*Token[T], error) {
	tokenParsed, err := ParseToken[T](token, secret)
	if err != nil {
		return nil, err
	}

	err = VerifyToken(tokenParsed, tokenType)
	if err != nil {
		return nil, err
	}

	return tokenParsed, nil
}
