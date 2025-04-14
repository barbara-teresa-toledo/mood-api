package handler

import (
	"encoding/json"
	"mood-api/internal/service"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if err := CreateUser(&user); err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	user, err := GetUserByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := api.tokenAuth.Encode(map[string]interface{}{
		"user_id": user.ID,
		"role":    user.Role,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
