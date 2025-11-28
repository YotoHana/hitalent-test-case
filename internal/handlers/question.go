package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/YotoHana/hitalent-test-case/internal/service"
)

type QuestionHandler struct {
	service service.QuestionService
}

func (h *QuestionHandler) Questions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getQuestions(w, r)
	case http.MethodPost:
		h.createQuestion(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *QuestionHandler) QuestionsID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getQuestionByID(w, r)
	case http.MethodDelete:
		h.deleteQuestionByID(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *QuestionHandler) getQuestions(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAllQuestions(r.Context())
	if err != nil {
		http.Error(w, "Intenal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *QuestionHandler) createQuestion(w http.ResponseWriter, r *http.Request) {
	var req models.CreateQuestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateQuestion(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h *QuestionHandler) getQuestionByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	id, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetQuestionByID(r.Context(), id)
	if err != nil {
		if result == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *QuestionHandler) deleteQuestionByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	id, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteQuestionByID(r.Context(), id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewQuestionHandler(service service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		service: service,
	}
}