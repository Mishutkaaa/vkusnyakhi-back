package models

type Food struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Image    *string `json:"image,omitempty"`
	Category *[]int  `json:"category,omitempty"`
	Brand    *string `json:"brand,omitempty"`
}
