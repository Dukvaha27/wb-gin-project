package main

import (
	"fmt"
	"net/http"
	"strconv"
	"wb-gin-project/internal/config"
	"wb-gin-project/internal/models"
	"wb-gin-project/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// товары
	router.GET("/products", func(ctx *gin.Context) {
		products := services.GetAllProducts()

		status := config.Ternary(products != nil, http.StatusOK, http.StatusBadRequest)
		data := config.Ternary(products != nil, gin.H{"data": products}, gin.H{"error": "Product is not found"})

		ctx.JSON(status, data)
	})

	router.GET("/products/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id must be a number"})
			return
		}

		product := services.GetProductById(id)

		status := config.Ternary(product != nil, http.StatusOK, http.StatusBadRequest)
		data := config.Ternary(product != nil, gin.H{"data": product}, gin.H{"error": "Product is not found"})

		ctx.JSON(status, data)
	})

	router.PATCH("/products/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id must be a number"})
			return
		}
		var req models.ProductCreate

		reqErr := ctx.ShouldBindJSON(&req)

		if reqErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error validation"})
			return
		}

		product := services.UpdateProduct(id, req)

		status := config.Ternary(product, http.StatusOK, http.StatusBadRequest)
		data := config.Ternary(product, gin.H{"data": "Продук успешно обновлен"}, gin.H{"error": "Product is not found"})

		ctx.JSON(status, data)
	})
	router.POST("/products", func(ctx *gin.Context) {
		var req models.ProductCreate

		err := ctx.ShouldBindJSON(&req)
		product := services.CreateProduct(req)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	})
	router.DELETE("/products/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "ID должен быть числом",
			})
			return
		}

		isDeleted := services.RemoveProduct(id, false)

		if isDeleted {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Товар успешно удален",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Товар с id=%d не найден", id),
			})
		}
	})

	// категории
	router.GET("/categories", func(ctx *gin.Context) {
		list := config.LoadCategories()

		ctx.JSON(http.StatusOK, list)
	})
	router.GET("/categories/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "id must be number",
			})
			return
		}
		category := services.GetCategoryById(id)

		if category == nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("category with id=%d is not found", id),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"data": category,
			})
		}
	})
	router.DELETE("/categories/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "id must be number",
			})
			return
		}

		isDeleted := services.RemoveCatregory(id)

		status := config.Ternary(isDeleted, http.StatusOK, http.StatusBadRequest)
		message := config.Ternary(isDeleted, gin.H{"message": "Category is deleted successful"}, gin.H{"message": fmt.Sprintf("category with id=%d is not found", id)})

		ctx.JSON(status, message)
	})
	router.POST("/categories", func(ctx *gin.Context) {
		var req models.CategoryPOST

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": services.CreateCategory(req)})
	})

	router.PATCH(
		"/categories/:id",
		func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id must be a number"})
				return
			}
			var req models.CategoryPOST

			reqErr := ctx.ShouldBindJSON(&req)

			if reqErr != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error validation"})
				return
			}

			category := services.UpdateCategory(id, req)

			status := config.Ternary(category, http.StatusOK, http.StatusBadRequest)
			data := config.Ternary(category, gin.H{"data": "Category успешно обновлен"}, gin.H{"error": "category is not found"})

			ctx.JSON(status, data)
		},
	)

	//комментарии
	router.GET("/comments", func(ctx *gin.Context) {
		list := config.LoadComments()
		ctx.JSON(http.StatusOK, list)
	})
	router.GET("/comments/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "id must be a string"})
			return
		}

		comment := services.GetCommentById(id)

		status := config.Ternary(comment != nil, http.StatusOK, http.StatusBadRequest)
		data := config.Ternary(comment != nil, gin.H{"data": comment}, gin.H{"message": "comment is not found"})

		ctx.JSON(status, data)
	})

	router.DELETE("/comments/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "id must be a string"})
			return
		}

		isDeleted := services.RemoveComment(id, false)

		status := config.Ternary(isDeleted, http.StatusOK, http.StatusBadRequest)
		data := config.Ternary(isDeleted, gin.H{"message": "Commen is deleted succesfull"}, gin.H{"message": "comment is not found"})

		ctx.JSON(status, data)
	})
	router.POST("/comments", func(ctx *gin.Context) {

		var req models.CommentPOST

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": services.Create(req)})
	})

	router.PATCH(
		"/comments/:id",
		func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id must be a number"})
				return
			}
			var req models.CommentPOST

			reqErr := ctx.ShouldBindJSON(&req)

			if reqErr != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error validation"})
				return
			}

			comment := services.UpdateComment(id, req)

			status := config.Ternary(comment, http.StatusOK, http.StatusBadRequest)
			data := config.Ternary(comment, gin.H{"data": "The comment успешно обновлен"}, gin.H{"error": "the comment is not found"})

			ctx.JSON(status, data)
		},
	)

	router.Run(":8080")
}
