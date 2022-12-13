package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Testing go and redis")
	port := "80"

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.GET("/response", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "A message from Redis: %s\n", pong)
	})

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
