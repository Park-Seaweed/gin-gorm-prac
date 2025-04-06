package dto

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=5"`
	Content string `json:"content" binding:"required,min=5,max=500"`

	//UserID uint `json:"-"`
}
