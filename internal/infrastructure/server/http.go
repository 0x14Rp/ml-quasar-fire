package server

import (
	"encoding/json"
	"fmt"
	"log"
	"ml-quasar-fire/internal/entity"
	"ml-quasar-fire/internal/usecase"
	"ml-quasar-fire/pkg/errs"
	"net/http"
)

var satelliteStore = make(map[string]entity.SatelliteData)

var validSatellites = map[string]bool{
	"kenobi":    true,
	"skywalker": true,
	"sato":      true,
}

func TopSecretHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errs.WriteJSONError(w, r, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	var req entity.TopSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "Cuerpo de la peticion Invalido")
		return
	}

	if len(req.Satellites) != 3 {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "Numero de satellites insuficientes")
		return
	}

	var messages [][]string
	var distances []float32

	for _, satellite := range req.Satellites {
		messages = append(messages, satellite.Message)
		distances = append(distances, satellite.Distance)
	}

	position, message, err := usecase.ProcessTopSecretRequest(distances, messages)
	if err != nil {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "insuficiente información para reconstruir el mensaje")
		return
	}
	resp := entity.TopSecretResponse{
		Position: position,
		Message:  message,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		errs.WriteJSONError(w, r, http.StatusMethodNotAllowed, "Error en la respuesta")
		return
	}
}

func TopSecretSplitHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		PostTopSecretSplitHandler(w, r)
	case http.MethodGet:
		GetTopSecretSplitHandler(w, r)
	case http.MethodDelete:
		ClearSatelliteStore(w, r)
	default:
		errs.WriteJSONError(w, r, http.StatusMethodNotAllowed, "Metodo no permitido")
	}
}

func PostTopSecretSplitHandler(w http.ResponseWriter, r *http.Request) {

	satelliteName := r.URL.Path[len("/topsecret_split/"):]

	if _, ok := validSatellites[satelliteName]; !ok {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "Satellite desconocido")
		return
	}

	var data entity.SatelliteData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "Cuerpo de la peticion Invalido")
		return
	}

	satelliteStore[satelliteName] = data

	message := fmt.Sprintf("Mensaje de %s guardado", satelliteName)
	response := map[string]interface{}{"message": message, "code": http.StatusOK}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "Error writing response")
		return
	}

}

func GetTopSecretSplitHandler(w http.ResponseWriter, r *http.Request) {
	if len(satelliteStore) < 3 {
		errs.WriteJSONError(w, r, http.StatusNotFound, "No hay suficiente información")
		return
	}

	var distances []float32
	var messages [][]string

	for _, data := range satelliteStore {
		distances = append(distances, data.Distance)
		messages = append(messages, data.Message)
	}

	x, y, err := usecase.GetLocation(distances...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	completeMessage, err := usecase.ProcessMessage(messages...)
	if err != nil {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "insuficiente información para reconstruir el mensaje")
		return
	}

	res := entity.Position{X: x, Y: y}
	resp := entity.TopSecretResponse{
		Position: res,
		Message:  completeMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		errs.WriteJSONError(w, r, http.StatusMethodNotAllowed, "Error en la respuesta")
	}
}

func ClearSatelliteStore(w http.ResponseWriter, r *http.Request) {
	if len(satelliteStore) > 0 {
		for k := range satelliteStore {
			delete(satelliteStore, k)
		}
		response := map[string]interface{}{"message": "Mensajes eliminados", "code": http.StatusOK}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	} else {
		errs.WriteJSONError(w, r, http.StatusBadRequest, "No existen mensajes")
	}
}
