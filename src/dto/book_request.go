package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required,max=200"`
	Author string `json:"author" validate:"omitempty,max=200"`
	Year   int    `json:"year"`
}
