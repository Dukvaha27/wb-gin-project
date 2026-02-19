package config

import "wb-gin-project/internal/models"

var products []models.Product = []models.Product{}
var categories = []models.Category{}
var comments = []models.Comment{}

func LoadProducts() []models.Product {
	result := make([]models.Product, len(products))
	copy(result, products)
	return result
}

func Save(p []models.Product) {
	products = make([]models.Product, len(p))
	copy(products, p)
}

func LoadCategories() []models.Category {
	result := make([]models.Category, len(categories))
	copy(result, categories)
	return result
}

func SaveCategories(p []models.Category) {
	categories = make([]models.Category, len(p))
	copy(categories, p)
}

func LoadComments() []models.Comment {
	result := make([]models.Comment, len(comments))
	copy(result, comments)
	return result
}

func SaveComments(c []models.Comment) {
	comments = make([]models.Comment, len(c))
	copy(comments, c)
}

func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}
