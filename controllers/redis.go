package controllers

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.com/alienate/password-generator/inter"
)

//HSET
//- HSet("myhash", "key1", "value1", "key2", "value2")
//- HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//- HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})

var redisCtx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "192.168.0.32:6379",
	Password: "n4th4n43l", // no password set
	DB:       0,           // use default DB
})

func Save(value *inter.NewPasswordReqSave) (string, error) {

	key := (uuid.New()).String()

	// Set some fields.
	_, err := rdb.Pipelined(redisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(redisCtx, key, "Account", "facebook")
		rdb.HSet(redisCtx, key, "Username", "nate@gmail.com")
		rdb.HSet(redisCtx, key, "Password", "n4th4n43l")
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil

}
func GetAll() (string, error) {

	var allFields inter.SavedFields
	// Scan all fields into the model.

	// Not sure about the star
	err := rdb.HGetAll(redisCtx, "*").Scan(&allFields)
	if err != nil {
		return "nil", err
	}

	return "success", nil

}
func Update(value *inter.SavedFields) (string, error) {

	// Set some fields.
	_, err := rdb.Pipelined(redisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(redisCtx, value.Key, "Account", value.Account)
		rdb.HSet(redisCtx, value.Key, "Username", value.Username)
		rdb.HSet(redisCtx, value.Key, "Password", value.Password)
		return nil
	})

	if err != nil {
		return "nil", err
	}

	return "success", nil
}
func Delete() {

}
