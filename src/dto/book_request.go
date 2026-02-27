package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required,max=200"`
	Author string `json:"author" validate:"omitempty,max=200"`
	Year   int    `json:"year"`
}

type BookQueryParam struct {
	Author string `query:"author"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}
