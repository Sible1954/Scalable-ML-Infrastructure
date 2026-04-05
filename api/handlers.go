
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scalable-ml-infrastructure/models"
)

func RegisterHandlers(router *http.ServeMux) {
	router.HandleFunc("/predict", handlePredict)
	router.HandleFunc("/health", handleHealthCheck)
	router.HandleFunc("/info", handleModelInfo)
}

func handlePredict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.PredictionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	prediction := models.PredictionResponse{
		ModelName:    "dummy-model-v1.0",
		PredictionID: "pred_12345",
		Result:       fmt.Sprintf("Prediction for features: %v", req.Features),
		Confidence:   0.95,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(prediction); err != nil {
		fmt.Printf("Error encoding prediction response: %v
", err)
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handleModelInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	info := models.ModelInfo{
		Name:        "dummy-model-v1.0",
		Version:     "1.0.0",
		Description: "A dummy model for demonstration purposes.",
		InputType:   "[]float64",
		OutputType:  "string",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}
