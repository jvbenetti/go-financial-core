package model

import "time"

type IdempotencyKey struct {
	Key          string    `json:"key"`
	ResponseBody []byte    `json:"response_body"` // Get the first response
	StatusCode   int       `json:"status_code"`
	CreatedAt    time.Time `json:"created_at"`
}
