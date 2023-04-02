package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// rds := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	DB:   0,
	// })
	// ctx := context.Background()
	// val, err := rds.Get(ctx, "key").Result()
	// if err != nil {
	// 	fmt.Println("No value of `key`")
	// }
	// fmt.Println(val)

	godotenv.Load()

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server: " + os.Getenv("PORT")))
	})

	server := http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("PORT"),
		Handler: handler,
	}

	fmt.Println("Server running on: " + "0.0.0.0:" + os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
