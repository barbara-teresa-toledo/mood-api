package internal

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func PatientOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["role"] != "paciente" {
			http.Error(w, "Acesso negado", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func PsychologistOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["role"] != "psicologo" {
			http.Error(w, "Acesso negado", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
