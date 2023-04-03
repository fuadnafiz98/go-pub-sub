package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	godotenv.Load()

	handler := http.NewServeMux()

	// TODO: make a redis service
	rds := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	ctx := context.Background()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rds.Publish(ctx, "messages", "Message from server: "+os.Getenv("PORT"))
		w.Write([]byte("Hello from server: " + os.Getenv("PORT")))
	})

	server := http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("PORT"),
		Handler: handler,
	}

	// on server initialization, connect to channel `messages`
	pubsub := rds.Subscribe(ctx, "messages")

	defer pubsub.Close()

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}

			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

	fmt.Println("Server running on: " + "0.0.0.0:" + os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
