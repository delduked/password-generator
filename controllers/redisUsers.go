package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/schema"
)

func CheckIfUserExists(value *schema.UserAccount) (bool, error) {

	var savedUser schema.UserAccount

	length, err := config.UsersRdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return false, err
	}

	for _, j := range length {
		err := config.UsersRdb.HGetAll(config.RedisCtx, j).Scan(&savedUser)
		if err != nil {
			return false, err
		}
		if savedUser.Username == value.Username && savedUser.Password == value.Password {
			return true, nil
		}
	}

	return false, nil
}

func SaveUser(value *schema.UserAccount) error {

	key := (uuid.New()).String()

	_, err := config.UsersRdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
