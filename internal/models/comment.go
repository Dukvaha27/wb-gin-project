package models

type Comment struct {
	ID        int    `json:"id"`
	Message   string `json:"message" binding:"required"`
	ProductId int    `json:"productId" binding:"required"`
}

type CommentPOST struct {
	Message   *string `json:"message" binding:"required"`
	ProductId *int    `json:"productId" binding:"required"`
}
