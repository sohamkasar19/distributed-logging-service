package logging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateLogHandler(t *testing.T) {
	log := Log{
		ID:        "test",
		Source:    "source",
		Message:   "message",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(log)

	req, err := http.NewRequest("POST", "/logs", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateLogHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if _, exists := Logs[log.ID]; !exists {
		t.Errorf("Log was not added")
	}
}
