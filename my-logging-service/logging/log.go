package logging

import "time"

type Log struct {
    ID        string    `json:"id"`
    Source    string    `json:"source"`
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}
