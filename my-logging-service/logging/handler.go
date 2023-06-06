package logging

import (
	"encoding/json"
	"net/http"
	"strings"
)

func LogsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllLogsHandler(w, r)
	case http.MethodPost:
		CreateLogHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func LogRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetLogHandler(w, r)
	case http.MethodPut:
		UpdateLogHandler(w, r)
	case http.MethodDelete:
		DeleteLogHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}



func CreateLogHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var newLog Log
    err := json.NewDecoder(r.Body).Decode(&newLog)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    go addLog(newLog)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newLog)
}

func GetAllLogsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    logs := getAllLogs()
    json.NewEncoder(w).Encode(logs)
}

func GetLogHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := strings.TrimPrefix(r.URL.Path, "/logs/")
    log, found := getLog(id)
    if found {
        json.NewEncoder(w).Encode(log)
    } else {
        http.Error(w, "Log not found", http.StatusNotFound)
    }
      
}

func UpdateLogHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := strings.TrimPrefix(r.URL.Path, "/logs/")
    var updatedLog Log
    err := json.NewDecoder(r.Body).Decode(&updatedLog)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updateLog(id, updatedLog)
    json.NewEncoder(w).Encode(updatedLog)
}

func DeleteLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := strings.TrimPrefix(r.URL.Path, "/logs/")
    deleteLog(id)
    w.WriteHeader(http.StatusNoContent)
}

func GetAggregatedLogsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    aggregatedLogs := getAggregatedLogs()
    json.NewEncoder(w).Encode(aggregatedLogs)
}
