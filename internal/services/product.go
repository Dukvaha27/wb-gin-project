package services

import (
	"wb-gin-project/internal/config"
	"wb-gin-project/internal/models"
)

func hasCategory(id int) bool {
	list := config.LoadCategories()

	for _, n := range list {
		if n.ID == id {
			return true
		}
	}
	return false
}

var productId = 0

func CreateProduct(p models.ProductCreate) models.Product {

	list := config.LoadProducts()
	product := models.Product{
		ID: productId,
	}

	if p.Name != nil {
		product.Name = *p.Name
	}
	if p.CategoryId != nil && hasCategory(*p.CategoryId) {
		product.CategoryId = *p.CategoryId
	}

	if p.Left != nil {
		product.Left = *p.Left
	}

	if p.IsVIP != nil {
		product.IsVIP = *p.IsVIP
	}

	if p.Price != nil {
		product.IsVIP = *p.IsVIP
	}

	list = append(list, product)
	config.Save(list)

	productId++
	return product
}

func UpdateProduct(id int, p models.ProductCreate) bool {
	list := config.LoadProducts()

	for idx, item := range list {
		if id == item.ID {
			currItem := &list[idx]
			if p.CategoryId != nil {
				currItem.CategoryId = *p.CategoryId
			}

			if p.IsVIP != nil {
				currItem.IsVIP = *p.IsVIP
			}

			if p.Left != nil {
				currItem.Left = *p.Left
			}

			if p.Name != nil {
				currItem.Name = *p.Name
			}

			if p.Price != nil {
				currItem.Price = *p.Price
			}

			config.Save(list)
			return true
		}
	}

	return false
}

func RemoveProduct(id int, cascade bool) bool {

	list := config.LoadProducts()

	change := false
	for i := 0; i < len(list); i++ {

		ID := config.Ternary(cascade, list[i].CategoryId, list[i].ID)

		if ID == id {
			change = true
			list = append(list[:i], list[i+1:]...)
			RemoveComment(id, true)
		}
	}

	if change {
		config.Save(list)
		return true
	}

	return false
}

func GetProductById(id int) *models.ProductGET {
	list := config.LoadProducts()
	comments := config.LoadComments()
	categories := config.LoadCategories()

	currComments := []models.Comment{}
	categoryMap := map[int]models.Category{}

	for _, n := range comments {
		if n.ProductId == id {
			currComments = append(currComments, n)
		}
	}

	for _, n := range categories {
		categoryMap[n.ID] = n
	}

	for _, n := range list {
		if n.ID == id {

			currCategory := categoryMap[n.CategoryId]
			return &models.ProductGET{
				ID:       n.ID,
				Price:    n.Price,
				Name:     n.Name,
				IsVIP:    n.IsVIP,
				Left:     n.Left,
				Category: currCategory.Name,
				Comments: currComments,
			}
		}
	}

	return nil
}

func GetAllProducts() *[]models.ProductGET {
	newList := []models.ProductGET{}

	list := config.LoadProducts()

	for _, n := range list {
		item := *GetProductById(n.ID)

		newList = append(newList, item)
	}

	return &newList
}
