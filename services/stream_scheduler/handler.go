package main

import (
	"log"

	"github.com/redis/go-redis/v9"
)

type handler struct {
	rdb *redis.Client
}

type Handler interface {
	Process(redisSortedSetKey string) error
}

func (h *handler) Process(redisSortedSetKey string) error {
	log.Printf("Successfully Processing %s", redisSortedSetKey)
	return nil
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
