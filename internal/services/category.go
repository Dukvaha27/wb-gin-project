package services

import (
	"wb-gin-project/internal/config"
	"wb-gin-project/internal/models"
)

var categoryId = 0

func CreateCategory(c models.CategoryPOST) models.Category {
	list := config.LoadCategories()
	category := models.Category{
		ID: categoryId,
	}

	if c.Name != nil {
		category.Name = *c.Name
	}

	categoryId++
	list = append(list, category)
	config.SaveCategories(list)

	return category
}

func UpdateCategory(id int, c models.CategoryPOST) bool {

	list := config.LoadCategories()

	for idx, n := range list {
		if n.ID == id && c.Name != nil {
			list[idx].Name = *c.Name

			config.SaveCategories(list)
			return true
		}
	}

	return false
}

func RemoveCatregory(id int) bool {

	list := config.LoadCategories()

	for i := 0; i < len(list); i++ {
		if list[i].ID == id {

			list = append(list[:i], list[i+1:]...)

			config.SaveCategories(list)
			RemoveProduct(id, true)
			return true
		}
	}

	return false
}

func GetCategoryById(id int) *models.Category {
	list := config.LoadCategories()

	for _, n := range list {

		if n.ID == id {
			return &n
		}
	}
	return nil
}
