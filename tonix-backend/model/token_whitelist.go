package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const TOKEN_WHITELIST_PREFIX = "token-whitelist"
var ctx = context.Background()

type TokenWhitelistFeatures struct {
	redisConn *redis.Client
}

func TokenWhitelist(redisConn *redis.Client) *TokenWhitelistFeatures {
	return &TokenWhitelistFeatures{
		redisConn: redisConn,
	}
}

func (t *TokenWhitelistFeatures) Add(tokenId string, expiresIn time.Duration) error {
	return t.redisConn.SetEx(ctx, TOKEN_WHITELIST_PREFIX + tokenId, "1", expiresIn).Err()
}

func (t *TokenWhitelistFeatures) IsWhitelisted(tokenId string) (bool, error) {
	res, err := t.redisConn.Exists(ctx, TOKEN_WHITELIST_PREFIX + tokenId).Result()

	return res == 1, err
}

func (t *TokenWhitelistFeatures) Remove(tokenId string) error {
	return t.redisConn.Del(ctx, TOKEN_WHITELIST_PREFIX + tokenId).Err()
}
