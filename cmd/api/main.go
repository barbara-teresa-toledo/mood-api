package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/lib/pq"

	"mood-api/internal/handler"
	"mood-api/internal/repository"
	"mood-api/internal/service"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()

	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

	userRepo := &repository.UserRepository{DB: db}
	moodRepo := &repository.MoodRepository{DB: db}

	authService := &service.AuthService{UserRepo: userRepo}
	moodService := &service.MoodService{MoodRepo: moodRepo}

	authHandler := &handler.AuthHandler{AuthService: authService, TokenAuth: tokenAuth}
	moodHandler := &handler.MoodHandler{MoodService: moodService}

	r := chi.NewRouter()

	r.Post("/signup", authHandler.SignUp)
	r.Post("/login", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/mood", moodHandler.CreateMoodEntry)
		r.Get("/report", moodHandler.GetMoodReport)
	})

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
