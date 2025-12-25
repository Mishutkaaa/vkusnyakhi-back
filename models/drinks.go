package models

import "github.com/lib/pq"

type Drinks struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Image      *string       `json:"image,omitempty"`
	Categories pq.Int32Array `json:"categories" db:"categories"`
	Brand      *int          `json:"brand,omitempty"`
}
