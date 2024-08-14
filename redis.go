package battlesshiplib

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedis() *Redis {
	return &Redis{
		rdb: Connect(),
		ctx: context.Background(),
	}
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

func (r Redis) Set(key string, value any, expiration time.Duration) error {
	_, err := r.rdb.Set(r.ctx, key, value, expiration).Result()
	if err != nil {
		fmt.Println("Can't set key:", err)
		return err
	}
	return nil
}

func (r Redis) Ping() {
	pong, err := r.rdb.Ping(r.ctx).Result()
	if err != nil {
		fmt.Println("Cannot connect to redis:", err)
		return
	}
	fmt.Println("Redis client started:", pong)
}

func (r Redis) Get(key string) (string, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("La clave no existe")
		} else {
			fmt.Println("No se pudo obtener la clave:", err)
		}
		return val, err
	}
	return val, nil
}

func (r Redis) Del(key string, value string) {
	err := r.rdb.Del(r.ctx, "key").Err()
	if err != nil {
		fmt.Println("No se pudo eliminar la clave:", err)
		return
	}
	fmt.Println("Clave eliminada")
}

func Encode[T any](value T) ([]byte, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode[T any](value []byte) (T, error) {
	var output T

	decoder := gob.NewDecoder(bytes.NewReader(value))
	err := decoder.Decode(&output)
	if err != nil {
		return output, err
	}
	return output, nil
}
