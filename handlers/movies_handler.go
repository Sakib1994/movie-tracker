package handlers

import (
	"encoding/json"
	"net/http"

	"sakib.com/reelingit/data"
	"sakib.com/reelingit/logger"
)

type MovieHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("JSON Encoding Error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}
func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetTopMovies()
	if h.handleStorageError(w, err, "Failed to get Top movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served top movies")
	}
}
func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetRandomMovies()
	if h.handleStorageError(w, err, "Failed to get random movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.logger.Info("Successfully served random movies")
	}
}
func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}
func NewMovieHandler(storage data.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		storage: storage,
		logger:  log,
	}
}
