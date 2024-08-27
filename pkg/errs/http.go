package errs

import (
	"encoding/json"
	"ml-quasar-fire/internal/entity"
	"net/http"
	"time"
)

func WriteJSONError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	errorResponse := entity.ErrorResponse{
		Status:    statusCode,
		Message:   message,
		Path:      r.URL.Path,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		w.Write([]byte("Error encoding JSON response"))
	}
}
