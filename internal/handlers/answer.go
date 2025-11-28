package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	}
}

func (h *AnswerHandler) createAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pathValue := r.PathValue("id")
	questionID, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateAnswer(r.Context(), req, questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func NewAnswerHandler(service service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		service: service,
	}
}
