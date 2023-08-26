package httpadapter

import (
	"encoding/json"
	"go-code-snippets/internal/core/domain/dto"
	"go-code-snippets/internal/core/helper"
	"go-code-snippets/internal/ports"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	httpPort ports.HttpPort
}

func NewHandler(httpPort ports.HttpPort) *Handler {
	return &Handler{
		httpPort: httpPort,
	}
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	createBookDto := dto.CreateBookDto{}
	if err := json.NewDecoder(r.Body).Decode(&createBookDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "invalid request in request body"})
		return
	}

	bookId, err := h.httpPort.CreateBook(r.Context(), createBookDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(helper.Response{"success": true, "book_id": bookId})
}

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "book_id")

	book, err := h.httpPort.GetBookById(r.Context(), bookId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(helper.Response{"success": true, "book": book})
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.httpPort.GetAllBooks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(helper.Response{"success": true, "books": books})
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBookDto := dto.UpdateBookDto{}
	bookId := chi.URLParam(r, "book_id")

	if err := json.NewDecoder(r.Body).Decode(&updateBookDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "invalid request in request body"})
		return
	}

	updatedBookId, err := h.httpPort.UpdateBook(r.Context(), bookId, updateBookDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Response{"success": true, "updated_book_id": updatedBookId})
}

func (h *Handler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "book_id")
	deletedBookId, err := h.httpPort.DeleteBookById(r.Context(), bookId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Response{"success": false, "error": "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Response{"success": true, "deleted_book_id": deletedBookId})
}
