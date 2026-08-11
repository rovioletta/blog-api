package main

import (
	"log/slog"
	"net/http"
	"os"

	"rovioletta/blog-api/api/article"
	"rovioletta/blog-api/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("Starting the Blog app...")

	defer gracefulShutdown(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", slog.String("error", err.Error()))
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.InitDB(logger)
	defer db.CloseDB()

	articleAPI := article.InitArticleServer(logger, db)

	r.Post("/article", articleAPI.CreateArticle)

	if err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), r); err != nil {
		logger.Error("Unable to run the server", slog.String("error", err.Error()))
		return
	}
}

func gracefulShutdown(logger *slog.Logger) {
	logger.Info("Closing the Blog app...")
}
