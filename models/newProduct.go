package models

type NewProduct struct {
	Name       string    `json:"name,omitempty"`
	Image      *string   `json:"image,omitempty"`
	Table      *string   `json:"table,omitempty"`
	Categories *[]string `json:"categories,omitempty"`
	Brand      *int      `json:"brand,omitempty"`
}
