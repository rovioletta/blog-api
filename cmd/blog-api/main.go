package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	article_api "rovioletta/blog-api/api/article"
	"rovioletta/blog-api/internal/article"
	database "rovioletta/blog-api/internal/db"
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

	db := database.NewDB(logger)
	defer db.CloseDB()

	articleServer := article.NewArticleService(db)
	articleAPI := article_api.NewArticleAPI(logger, articleServer)

	r.Post("/article", articleAPI.CreateArticle)
	r.Post("/article/list", articleAPI.GetArticlesByFilter)
	r.Get("/article/{id}", articleAPI.GetArticleByID)
	r.Delete("/article/{id}", articleAPI.DeleteArticleByID)

	if err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), r); err != nil {
		logger.Error("Unable to run the server", slog.String("error", err.Error()))
		return
	}
}

func gracefulShutdown(logger *slog.Logger) {
	logger.Info("Closing the Blog app...")
}
