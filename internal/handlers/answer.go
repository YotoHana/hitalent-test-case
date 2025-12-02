package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/YotoHana/hitalent-test-case/internal/service"
)

type AnswerHandler struct {
	service service.AnswerService
}

func (h *AnswerHandler) QuestionsIDAnswers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createAnswer(w, r)
	default:
		http.Error(w, "Method nod allowed", http.StatusMethodNotAllowed)
		log.Printf("URL: %s, method: %s, status: %d", r.URL.Path, r.Method, http.StatusMethodNotAllowed)
	}
}

func (h *AnswerHandler) AnswersID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAnswer(w, r)
	case http.MethodDelete:
		h.deleteAnswer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("URL: %s, method: %s, status: %d", r.URL.Path, r.Method, http.StatusMethodNotAllowed)
	}
}

func (h *AnswerHandler) createAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAnswerRequest

	start := time.Now()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusBadRequest, time.Since(start), err)
		return
	}

	pathValue := r.PathValue("id")
	questionID, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusBadRequest, time.Since(start), err)
		return
	}

	result, err := h.service.CreateAnswer(r.Context(), req, questionID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusInternalServerError, time.Since(start), err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
	log.Printf("URL: %s, method: %s, status: %d, duration: %v", r.URL.Path, r.Method, http.StatusCreated, time.Since(start))
}

func (h *AnswerHandler) getAnswer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	pathValue := r.PathValue("id")
	id, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusBadRequest, time.Since(start), err)
		return
	}

	result, err := h.service.GetAnswer(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusInternalServerError, time.Since(start), err)
		return
	}

	w.Header().Set("Content-Type", "apllication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	log.Printf("URL: %s, method: %s, status: %d, duration: %v", r.URL.Path, r.Method, http.StatusOK, time.Since(start))
}

func (h *AnswerHandler) deleteAnswer(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	pathValue := r.PathValue("id")
	id, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusInternalServerError, time.Since(start), err)
		return
	}

	if err = h.service.DeleteAnswer(r.Context(), id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("URL: %s, method: %s, status: %d, duration: %v, error: %v", r.URL.Path, r.Method, http.StatusInternalServerError, time.Since(start), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("URL: %s, method: %s, status: %d, duration: %v", r.URL.Path, r.Method, http.StatusNoContent, time.Since(start))
}

func NewAnswerHandler(service service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		service: service,
	}
}
