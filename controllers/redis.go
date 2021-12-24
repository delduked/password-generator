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
	//	var keyedField []types.KeyedField
	savedFields := []types.SavedFields{}
	var keyedField types.KeyedField

	// Not sure about the star
	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return savedFields, err
	}

	for i, j := range length {
		err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
		if err != nil {
			fmt.Println(keyedField)
			return savedFields, err
		}
		fmt.Println(i)
		savedField := types.SavedFields{
			Key:      j,
			Account:  keyedField.Account,
			Username: keyedField.Username,
			Password: keyedField.Password,
		}
		fmt.Println(keyedField)
		savedFields = append(savedFields, savedField)
	}

	fmt.Println(length)

	return savedFields, err

}
func GetKeyedPassword(key string) (types.KeyedField, error) {

	var keyedField types.KeyedField

	// Not sure the key
	err := config.Rdb.HGetAll(config.RedisCtx, key).Scan(&keyedField)
	//err := config.Rdb.HMGet(config.RedisCtx, key).Scan(keyedField)
	//err := config.Rdb.HGetAll(config.RedisCtx, key).Scan(&KeyedField)
	if err != nil {
		fmt.Println(keyedField)
		return keyedField, err
	}
	fmt.Println(keyedField)
	return keyedField, nil

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
