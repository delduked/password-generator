package controllers

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/types"
)

func Save(value *types.NewPasswordReqSave) (string, error) {

	key := (uuid.New()).String()

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, key, "Password", value.Password)
		fmt.Println(value.Account)
		fmt.Println(value.Username)
		fmt.Println(value.Password)
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil

}
func GetAll() ([]types.SavedFields, error) {

	// Scan all fields into the model.
	var allFields []types.SavedFields

	// Not sure about the star
	err := config.Rdb.HGetAll(config.RedisCtx, "*").Scan(&allFields)
	if err != nil {
		fmt.Println(allFields)
		return allFields, err
	}

	return allFields, nil

}
func Update(value *types.SavedFields) (string, error) {

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
