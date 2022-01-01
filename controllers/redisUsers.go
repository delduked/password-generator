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

func SaveUser(value *schema.SignUp) error {

	// Generate a new uuid for the user if it does not exist
	exists, err := doesUserExist(value)
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
func doesUserExist(value *schema.SignUp) (bool, error) {

	var savedUser schema.UserAccount

	// Count how many fields exist in the database
	length, err := config.UsersRdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		// If there is an error checking the redis server, return the error
		return false, err
	}

	// Loop over all the fields saved in the database
	for _, j := range length {
		// Get each fields value using the uuid and parse into the struct
		err := config.UsersRdb.HGetAll(config.RedisCtx, j).Scan(&savedUser)
		if err != nil {
			// If there is an error, return it
			return false, err
		}
		if savedUser.Username == value.Username {
			// If the username exists, return true
			return true, nil
		}
	}
	// If the username does not exist, return false with no error
	return false, nil
}
