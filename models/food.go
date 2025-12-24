package models

type Food struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Image      *string `json:"image,omitempty"`
	Categories *[]int  `json:"categories,omitempty"`
	Brand      *int    `json:"brand,omitempty"`
}
