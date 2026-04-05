
package models

type PredictionRequest struct {
	Features []float64 `json:"features"`
}

type PredictionResponse struct {
	ModelName    string      `json:"model_name"`
	PredictionID string      `json:"prediction_id"`
	Result       interface{} `json:"result"`
	Confidence   float64     `json:"confidence"`
}

type ModelInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	InputType   string `json:"input_type"`
	OutputType  string `json:"output_type"`
}
