package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/mdayat/fullstack2024-test/go/configs"
)

func NewRestHandler(configs configs.Configs) *chi.Mux {
	router := chi.NewRouter()

	router.Use(chiMiddleware.CleanPath)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(httprate.LimitByIP(100, 1*time.Minute))

	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "Host", "Origin", "Referer", "Authorization"},
		ExposedHeaders:   []string{"Content-Length", "Location"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	router.Use(cors.Handler(options))
	router.Use(chiMiddleware.Heartbeat("/ping"))

	myClientHandler := NewMyClientHandler(configs)
	router.Post("/my-clients", myClientHandler.CreateMyClient)
	router.Put("/my-clients/{myClientId}", myClientHandler.UpdateMyClient)
	router.Delete("/my-clients/{myClientId}", myClientHandler.DeleteMyClient)

	return router
}
