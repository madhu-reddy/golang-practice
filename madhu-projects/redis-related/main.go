package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "mypassword", // no password set
		DB:       0,            // use default DB
	})

	err := rdb.Set(ctx, "ddd", "vae", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "ddd").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}
