package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database := db.Connect()
	defer database.Close()
	log.Println("Conexion a la base de datos exitosa")

	handlers.InitializeServices()
	log.Println("Servicios inicializados exitosamente")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/groups", handlers.GetGroups)
	r.Post("/users", handlers.CreateUser)
	r.Get("/users", handlers.GetUsers)
	r.Delete("/users/{userID}", handlers.DeleteUser)
	r.Post("/groups", handlers.CreateGroup)
	r.Delete("/groups/{groupID}", handlers.DeleteGroup)
	r.Post("/groups/{groupID}/join", handlers.JoinGroup)
	r.Delete("/groups/{groupID}/users/{userID}", handlers.LeftGroup)
	r.Patch("/groups/{groupID}/users/{userID}/role", handlers.ChangeRole)
	r.Post("/groups/{groupID}/categories", handlers.CreateCategory)
	r.Get("/groups/{groupID}/categories", handlers.GetCategoriesByGroup)
	r.Delete("/groups/{groupID}/categories/{categoryID}", handlers.DeleteCategory)
	r.Post("/groups/{groupID}/categories/{categoryID}/channels", handlers.CreateChannel)
	r.Get("/groups/{groupID}/channels", handlers.GetChannelsByGroup)
	r.Get("/groups/{groupID}/categories/{categoryID}/channels", handlers.GetChannelByCategory)
	r.Get("/groups/{groupID}/users", handlers.GetAllUsersByGroup)

	r.Delete("/groups/{groupID}/categories/{categoryID}/channels/{channelID}", handlers.DeleteChannel)
	log.Println("Starting server on http://localhost:8080")
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API corriendo 🚀"))
	})

	http.ListenAndServe(":8080", corsMiddleware(r))

}

func corsMiddleware(next http.Handler) http.Handler {
	allowedOrigin := "https://sync-flow-groups.vercel.app/"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
