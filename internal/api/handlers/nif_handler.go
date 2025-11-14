package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Etevaldo15/go-nif-validator/internal/core/nif"
)

// Response representa a resposta da validação.
type Response struct {
	NIF     string `json:"nif"`
	IsValid bool   `json:"is_valid"`
	Type    string `json:"type,omitempty"` // pode ser vazio se inválido
	Message string `json:"message"`
}

// ValidateNIF valida um NIF angolano.
//
// @Summary Valida um NIF angolano
// @Description Verifica se o NIF informado corresponde ao formato válido de Pessoa Singular (14 caracteres) ou Pessoa Coletiva (10 dígitos).
// @Tags nif
// @Accept  json
// @Produce  json
// @Param nif path string true "NIF a ser validado" minlength(1)
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/v1/validate-nif/{nif} [get]
func ValidateNIF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nifParam := vars["nif"]

	if nifParam == "" {
		sendJSONError(w, http.StatusBadRequest, "NIF é obrigatório", "")
		return
	}

	// Detecta o tipo **antes** da validação (não depende do resultado dela)
	nifType := nif.DetectType(nifParam)

	// Valida apenas o formato
	isValid, message := nif.Validate(nifParam)

	resp := Response{
		NIF:     nifParam,
		IsValid: isValid,
		Type:    string(nifType), // pode ser "Desconhecido" se formato inválido
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resp)
}

func sendJSONError(w http.ResponseWriter, status int, message, nif string) {
	resp := Response{
		NIF:     nif,
		IsValid: false,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}