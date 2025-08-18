package model

import "time"

type Entry struct {
	Name         string    `json:"name"`
	Host         string    `json:"host"`
	User         string    `json:"user"`
	Port         int       `json:"port"`
	KeyPath      string    `json:"keyPath,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	AttachmentID string    `json:"attachmentId,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
}
