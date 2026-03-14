// Package models
package models

import "time"

type ShortenedURL struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
