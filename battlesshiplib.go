package battlesshiplib

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb    *redis.Client
	ctx    context.Context
	stream bytes.Buffer
	dec    *gob.Decoder
	enc    *gob.Encoder
}

func NewRedis() *Redis {
	s := bytes.Buffer{}
	return &Redis{
		rdb:    Connect(),
		ctx:    context.Background(),
		stream: s,
		dec:    gob.NewDecoder(&s),
		enc:    gob.NewEncoder(&s),
	}
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//defer rdb.Close()

	return rdb
}

func (r Redis) Set(key string, value any) {
	_, err := r.rdb.Set(r.ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println("Can't set key:", err)
		return
	}
}

func (r Redis) Ping() {
	pong, err := r.rdb.Ping(r.ctx).Result()
	if err != nil {
		fmt.Println("Cannot connect to redis:", err)
		return
	}
	fmt.Println("Redis client started:", pong)
}

func (r Redis) Get(key string) any {
	val, err := r.rdb.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("La clave no existe")
		} else {
			fmt.Println("No se pudo obtener la clave:", err)
		}
		return nil
	}
	return val
}

func (r Redis) Del(key string, value string) {
	err := r.rdb.Del(r.ctx, "key").Err()
	if err != nil {
		fmt.Println("No se pudo eliminar la clave:", err)
		return
	}
	fmt.Println("Clave eliminada")
}

func (r Redis) Encode(value any) ([]byte, error) {
	var buf bytes.Buffer
	err := r.enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (r Redis) Decode(value []byte) (any, error) {
	var output any
	err := r.dec.Decode(&output)
	if err != nil {
		return output, err
	}
	return output, nil
}
