package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type CategoryPOST struct {
	Name *string
}
