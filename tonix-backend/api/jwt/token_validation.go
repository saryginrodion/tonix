package jwt

import (
	"strings"
	"time"
	"tonix/backend/api/utils"
)

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
	if token.Payload.Exp < time.Now().Unix() {
		return NewJWTError("JWT expired")
	}

	if token.Payload.TokenType != tokenType {
		return NewJWTError("JWT incorrect token type")
	}

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
