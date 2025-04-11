package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mdayat/fullstack2024-test/go/configs"
	"github.com/mdayat/fullstack2024-test/go/internal/handlers"
)

func main() {
	env, err := configs.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	db, err := configs.NewDb(ctx, env.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	redis := configs.NewRedis(env.RedisURL)
	defer redis.Close()

	configs := configs.NewConfigs(env, db, redis)
	router := handlers.NewRestHandler(configs)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
