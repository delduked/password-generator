package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/inter"
)

func Save(value *inter.NewPasswordReqSave) (string, error) {

	key := (uuid.New()).String()

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil

}
func GetAll() ([]inter.SavedFields, error) {

	var allFields []inter.SavedFields
	// Scan all fields into the model.

	// Not sure about the star
	err := config.Rdb.HGetAll(config.RedisCtx, "*").Scan(&allFields)
	if err != nil {
		return allFields, err
	}

	return allFields, nil

}
func Update(value *inter.SavedFields) (string, error) {

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, value.Key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, value.Key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, value.Key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil
}

// func Delete() {

// }
