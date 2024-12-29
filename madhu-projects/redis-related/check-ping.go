package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go Ping Test")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "mypassword",
		DB:       0,
	})

	// Ping test 1st way
	ping := client.Ping(context.Background())
	fmt.Println(ping)

	// Ping test 2nd way - best way
	ping2, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping2)
}

/*
Output

Go Ping Test
ping: PONG
PONG

*/
