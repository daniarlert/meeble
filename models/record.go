package models

import "time"

// A Record represents a new record of the user mood with an optional
// reason/explanation. It also keeps track of when the record was created.
type Record struct {
	Id        int64     `json:"id"`
	Mood      string    `json:"mood"`
	Reason    string    `json:"reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// NewRecord creates a new record with the received `mood` and `reason` arguments
func NewRecord(mood string, reason string) *Record {
	return &Record{
		Id:        42,
		Mood:      mood,
		Reason:    reason,
		CreatedAt: time.Now().UTC(),
	}
}
