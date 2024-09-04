package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"simple_http_server/models"
	"simple_http_server/services"
)

var (
	mutex sync.Mutex
)

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req models.Request
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.Cmd == "" || req.ID == 0 {
		http.Error(w, "Invalid 'cmd' or 'id' field", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Handle the request using the state service
	resp := services.HandleRequest(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
