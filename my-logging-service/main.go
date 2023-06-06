package main

import (
	"my-logging-service/logging"
	"net/http"
	"time"
)

func main() {
    go func() {
        for {
            logging.AggregateLogs()
            time.Sleep(1 * time.Minute)
        }
    }()

    http.HandleFunc("/logs", logging.LogsRouter)
    http.HandleFunc("/logs/", logging.LogRouter)
    http.HandleFunc("/aggregatedlogs", logging.GetAggregatedLogsHandler)
    
    http.ListenAndServe(":8080", nil)
}
