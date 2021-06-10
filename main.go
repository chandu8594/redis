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

	key, value := "Namae", "Uzumaki Naruto"
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("Error while setting name in redis: ", err.Error())
	}

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error while fetching name from redis: ", err.Error())
	}
	fmt.Println(key, " -> ", val)
}