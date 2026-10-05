package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type calculationRequest struct {
	A         *float64 `json:"a"`
	B         *float64 `json:"b"`
	Operation string   `json:"operacao"`
}

type calculationResponse struct {
	Operation string  `json:"operacao"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"resultado"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/calcular", calculate)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "servico": "calculadora-go"})
	})

	log.Println("API Go da calculadora ouvindo na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"erro": "Use o método POST."})
		return
	}

	var request calculationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil || request.A == nil || request.B == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"erro": "Informe dois números válidos em 'a' e 'b'."})
		return
	}

	a, b := *request.A, *request.B
	var result float64

	switch request.Operation {
	case "somar":
		result = a + b
	case "subtrair":
		result = a - b
	case "multiplicar":
		result = a * b
	case "dividir":
		if b == 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"erro": "Não é possível dividir por zero."})
			return
		}
		result = a / b
	default:
		writeJSON(w, http.StatusBadRequest, map[string]string{"erro": "Operação inválida."})
		return
	}

	writeJSON(w, http.StatusOK, calculationResponse{
		Operation: request.Operation,
		A:         a,
		B:         b,
		Result:    result,
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
