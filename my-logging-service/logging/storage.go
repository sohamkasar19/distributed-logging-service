package logging

import (
	"sync"
	"time"
)

// Logs is a map where the key is the log's ID and the value is the Log struct.
var (
    Logs           = make(map[string]Log)
    AggregatedLogs = make(map[string]Log)
    mu             = &sync.Mutex{}
)


// addLog adds a log to the Logs map.
func addLog(log Log) {
    mu.Lock()
    Logs[log.ID] = log
    mu.Unlock()
}

// getAllLogs returns a slice of all logs in the Logs map.
func getAllLogs() []Log {
    mu.Lock()
    logList := make([]Log, 0, len(Logs))
    for _, log := range Logs {
        logList = append(logList, log)
    }
    mu.Unlock()
    return logList
}

// getLog returns a log with a specific ID from the Logs map.
// The boolean value indicates whether the log was found.
func getLog(id string) (Log, bool) {
    mu.Lock()
    log, found := Logs[id]
    mu.Unlock()
    return log, found
}

// updateLog updates the log with a specific ID in the Logs map with new log data.
func updateLog(id string, newLog Log) {
    mu.Lock()
    Logs[id] = newLog
    mu.Unlock()
}

// deleteLog removes the log with a specific ID from the Logs map.
func deleteLog(id string) {
    mu.Lock()
    delete(Logs, id)
    mu.Unlock()
}

func AggregateLogs() {
    mu.Lock()
    defer mu.Unlock()

    // Clear the previous aggregated logs
    AggregatedLogs = make(map[string]Log)

    // Filter and aggregate logs from the past hour
    cutoff := time.Now().Add(-1 * time.Hour)
    for _ , log := range Logs {
        if log.Timestamp.After(cutoff) {
            // Aggregate logs by source
            aggLog, exists := AggregatedLogs[log.Source]
            if exists {
                aggLog.Message += "\n" + log.Message
            } else {
                aggLog = Log{Source: log.Source, Message: log.Message, Timestamp: log.Timestamp}
            }
            AggregatedLogs[log.Source] = aggLog
        }
    }
}

func getAggregatedLogs() map[string]Log {
    mu.Lock()
    aggregatedLogs := AggregatedLogs
    mu.Unlock()
    return aggregatedLogs
}

