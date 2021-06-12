package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	redisCache := NewRedisCache(client)
	ctx := context.Background()

	key, value := "Namae", "Uzumaki Naruto"
	if _, err := redisCache.Set(ctx, key, value, 0); err != nil {
		fmt.Println("Error while setting name in redis: ", err.Error())
	}

	val, _ := redisCache.Get(ctx, key)
	if val == nil {
		fmt.Println("Invalid key: ", key)
	} else {
		fmt.Println(key, " -> ", val)
	}

	if err := redisCache.Remove(ctx, key); err != nil {
		fmt.Printf("unable to remove key: %s from cache\n", key)
	}

	val, _ = redisCache.Get(ctx, key)
	if val == nil {
		fmt.Println("Invalid key: ", key)
	} else {
		fmt.Println(key, " -> ", val)
	}
}