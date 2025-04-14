package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"mood-api/internal/models"
	"mood-api/internal/service"
)

type MoodHandler struct {
	MoodService *service.MoodService
}

func (h *MoodHandler) CreateMoodEntry(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["user_id"].(string)

	var input struct {
		Level int `json:"level"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	entry := &models.MoodEntry{
		UserID: userID,
		Level:  input.Level,
	}

	if err := h.MoodService.AddMoodEntry(entry); err != nil {
		http.Error(w, "could not save mood entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MoodHandler) GetMoodReport(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID := claims["user_id"].(string)

	report, err := h.MoodService.GenerateReport(userID)
	if err != nil {
		http.Error(w, "could not generate report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
