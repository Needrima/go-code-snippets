package routes

import (
	httpadapter "go-code-snippets/internal/adapters/http-adapter"
	"go-code-snippets/internal/core/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(handler *httpadapter.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.JSONMiddleware)

	router.Post("/api/create-book", handler.CreateBook)
	router.Get("/api/get-book/{book_id}", handler.GetBookById)
	router.Get("/api/get-all-books", handler.GetAllBooks)
	router.Put("/api/update-book/{book_id}", handler.UpdateBook)
	router.Delete("/api/delete-book/{book_id}", handler.DeleteBookById)

	return router
}
