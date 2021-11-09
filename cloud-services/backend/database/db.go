package database

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

var db *redis.Client

func init() {
	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

type v struct {
	Vals []string `json:"vals"`
}

func GetFromDB(key string) ([]string, error) {
	b, err := db.Get(key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var arr v
	if err := json.Unmarshal(b, &arr); err != nil {
		return nil, err
	}
	return arr.Vals, nil
}

func AddToDB(key string, value string) error {
	b, err := db.Get(key).Bytes()
	if err != nil && err != redis.Nil {
		return err
	}
	var arr v
	if err != redis.Nil {
		if err := json.Unmarshal(b, &arr); err != nil {
			return err
		}
	}
	arr.Vals = append(arr.Vals, value)
	b, err = json.Marshal(arr)
	if err != nil {
		return err
	}
	err = db.Set(key, b, 0).Err()
	return err
}
