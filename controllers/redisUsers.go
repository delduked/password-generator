package controllers

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/schema"
)

func CheckIfOldUserExists(value *schema.UserAccount) (bool, error) {

	fields, err := config.UsersRdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return false, err
	}

	for _, j := range fields {
		field, err := config.UsersRdb.HGetAll(config.RedisCtx, j).Result()
		if err != nil {
			return false, err
		} else if len(field) == 0 {
			return false, err
		} else if err == redis.Nil {
			return false, err
		}

		if field["Username"] == value.Username && field["Password"] == value.Password {
			return true, nil
		}
	}

	return false, fmt.Errorf("Username or password is incorrect")
}

func SaveUser(value *schema.SignUp) error {

	// Generate a new uuid for the user if it does not exist
	exists, err := CheckIfNewUserExists(value)
	if err != nil || exists {
		return err
	}

	key := (uuid.New()).String()

	_, err = config.UsersRdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
func CheckIfNewUserExists(value *schema.SignUp) (bool, error) {

	//var keyedField schema.UserAccount
	//var savedField schema.SavedField

	length, err := config.UsersRdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return false, err
	}

	for _, j := range length {
		length, err := config.UsersRdb.HGetAll(config.RedisCtx, j).Result()
		if err != nil {
			return false, err
		} else if len(length) == 0 {
			return false, err
		} else if err == redis.Nil {
			return false, err
		}

		if length["Username"] == value.Username {
			return true, fmt.Errorf("Username already exists")
		}
	}
	// If the username does not exist, return false with no error
	return false, nil
}
