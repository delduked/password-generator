package controllers

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/config"
	"gitlab.com/alienate/password-generator/schema"
)

func SaveMany(username string, value []schema.KeyedField) error {

	var err error
	var key string
	for _, j := range value {
		key = (uuid.New()).String()
		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
			rdb.HSet(config.RedisCtx, username+"::"+key, "Account", j.Account)
			rdb.HSet(config.RedisCtx, username+"::"+key, "Username", j.Username)
			rdb.HSet(config.RedisCtx, username+"::"+key, "Password", j.Password)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return err
}

func Save(username string, value *schema.KeyedField) (schema.SavedField, error) {

	key := (uuid.New()).String()
	var savedField schema.SavedField

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, username+"::"+key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, username+"::"+key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, username+"::"+key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return savedField, err
	}

	savedField = schema.SavedField{
		Key:      username + "::" + key,
		Account:  value.Account,
		Username: value.Username,
		Password: value.Password,
	}

	return savedField, nil
}
func GetAll(username string) ([]schema.SavedField, error) {

	savedFields := []schema.SavedField{}

	length, err := config.Rdb.Keys(config.RedisCtx, username+"::*").Result()
	if err != nil {
		return savedFields, err
	}

	var savedField schema.SavedField
	for _, j := range length {
		field, err := config.Rdb.HGetAll(config.RedisCtx, j).Result()
		if err == nil {
			savedField = schema.SavedField{
				Key:      j,
				Account:  field["Account"],
				Username: field["Username"],
				Password: field["Password"],
			}
			savedFields = append(savedFields, savedField)
		}
	}

	return savedFields, err
}
func GetKeyedPassword(key string) (schema.KeyedField, error, int) {

	var keyedField schema.KeyedField
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
func Update(username string, value *schema.SavedField) (string, error) {

	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, username+"::"+value.Key, "Account", value.Account)
		rdb.HSet(config.RedisCtx, username+"::"+value.Key, "Username", value.Username)
		rdb.HSet(config.RedisCtx, username+"::"+value.Key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil
}

func Delete(key string) (schema.KeyedField, error) {

	var keyedField schema.KeyedField

	err := config.Rdb.Del(config.RedisCtx, key).Err()
	if err != nil {
		return keyedField, err
	}
	return keyedField, nil
}
