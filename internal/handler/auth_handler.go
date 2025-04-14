package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"mood-api/internal/models"
	"mood-api/internal/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
	TokenAuth   *jwtauth.JWTAuth
}

// POST /signup
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := h.AuthService.RegisterUser(&user); err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// POST /login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.TokenAuth.Encode(map[string]interface{}{
		"user_id": user.ID,
		"role":    user.Role,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
