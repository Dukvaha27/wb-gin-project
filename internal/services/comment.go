package services

import (
	"wb-gin-project/internal/config"
	"wb-gin-project/internal/models"
)

var commentId = 0

func hasProduct(id int) bool {

	list := config.LoadProducts()

	for _, n := range list {
		if n.ID == id {
			return true
		}
	}
	return false
}

func Create(c models.CommentPOST) models.Comment {

	list := config.LoadComments()
	comment := models.Comment{
		ID: commentId,
	}

	if c.Message != nil {
		comment.Message = *c.Message
	}

	if c.ProductId != nil && hasProduct(*c.ProductId) {
		comment.ProductId = *c.ProductId
	}

	list = append(list, comment)
	config.SaveComments(list)

	commentId++
	return comment
}

func UpdateComment(id int, c models.CommentPOST) bool {
	list := config.LoadComments()

	for idx, item := range list {
		if item.ID == id && item.ProductId == *c.ProductId {
			list[idx].Message = *c.Message

			config.SaveComments(list)
			return true
		}
	}

	return false
}

func RemoveComment(id int, isCascade bool) bool {
	list := config.LoadComments()

	changed := false
	for idx, item := range list {
		ID := config.Ternary(isCascade, item.ProductId, item.ID)

		if id == ID {
			list = append(list[:idx], list[idx+1:]...)
			config.SaveComments(list)
			changed = true
		}
	}

	if changed {
		return true
	}
	return false
}

func GetCommentById(id int) *models.Comment {
	list := config.LoadComments()

	for _, n := range list {
		if n.ID == id {
			return &n
		}
	}
	return nil
}
