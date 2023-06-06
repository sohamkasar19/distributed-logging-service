package logging

import (
	"testing"
	"time"
)

func TestAddLog(t *testing.T) {
    log := Log{
        ID:        "test",
        Source:    "source",
        Message:   "message",
        Timestamp: time.Now(),
    }
    addLog(log)

    if _, exists := Logs[log.ID]; !exists {
        t.Errorf("Log was not added")
    }
}
