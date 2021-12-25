package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/types"
)

func Save(value *types.NewPasswordReqSave) (types.SavedField, error) {

	key := (uuid.New()).String()
	var savedField types.SavedField

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return savedField, err
	}

	savedField = types.SavedField{
		Key:      key,
		Account:  value.Account,
		Username: value.Username,
		Password: value.Password,
	}

	return savedField, nil

}
func GetAll() ([]types.SavedField, error) {

	savedFields := []types.SavedField{}
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
		savedField := types.SavedField{
			Key:      j,
			Account:  keyedField.Account,
			Username: keyedField.Username,
			Password: keyedField.Password,
		}
		savedFields = append(savedFields, savedField)
	}

	return savedFields, err
}
func GetKeyedPassword(key string) (types.KeyedField, error, int) {

	var keyedField types.KeyedField
	var Error error
	length, err := config.Rdb.HGetAll(config.RedisCtx, key).Result()
	if err == redis.Nil {
		return keyedField, Error, len(length)
	} else if err != nil {
		return keyedField, Error, len(length)
	} else if len(length) == 0 {
		return keyedField, Error, len(length)
	}

	err = config.Rdb.HGetAll(config.RedisCtx, key).Scan(&keyedField)
	if err != nil {
		return keyedField, err, len(length)
	}
	return keyedField, nil, len(length)
}
func Update(value *types.SavedField) (string, error) {

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
