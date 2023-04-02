package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx := context.Background()
	val, err := rds.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("No value of `key`")
	}
	fmt.Println(val)
}
