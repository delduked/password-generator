package controllers

import (
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
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil

}
func GetAll() ([]types.SavedFields, error) {

	savedFields := []types.SavedFields{}
	var keyedField types.KeyedField

	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return savedFields, err
	}

	for _, j := range length {
		err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedField)
		if err != nil {
			return savedFields, err
		}
		savedField := types.SavedFields{
			Key:      j,
			Account:  keyedField.Account,
			Username: keyedField.Username,
			Password: keyedField.Password,
		}
		savedFields = append(savedFields, savedField)
	}

	return savedFields, err
}
func GetKeyedPassword(key string) (types.KeyedField, error) {

	var keyedField types.KeyedField

	err := config.Rdb.HGetAll(config.RedisCtx, key).Scan(&keyedField)
	if err != nil {
		return keyedField, err
	}
	return keyedField, nil
}
func Update(value *types.SavedFields) (string, error) {

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

func Delete(key string) (types.KeyedField, error) {

	var keyedField types.KeyedField

	err := config.Rdb.Del(config.RedisCtx, key).Err()
	//err := config.Rdb.HGetAll(config.RedisCtx, key).Scan(&keyedField)
	if err != nil {
		return keyedField, err
	}
	return keyedField, nil
}
