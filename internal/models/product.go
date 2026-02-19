package models

type Product struct {
	ID         int        `json:"id" binding:"required"`
	Name       string     `json:"name" binding:"required"`
	Price      string     `json:"price" binding:"required"`
	Left       int        `json:"left" binding:"required"`
	IsVIP      bool       `json:"isVIP" binding:"required"`
	CategoryId int        `json:"categoryId" binding:"required"`
}

type ProductGET struct {
	ID       int       `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Price    string    `json:"price" binding:"required"`
	Left     int       `json:"left" binding:"required"`
	IsVIP    bool      `json:"isVIP" binding:"required"`
	Category string    `json:"category" binding:"required"`
	Comments []Comment `json:"comments" binding:"required"`
}


type ProductCreate struct {
	Name       *string
	Price      *string
	Left       *int
	IsVIP      *bool
	CategoryId *int
}