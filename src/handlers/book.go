package handlers

import (
	"desent/src/dto"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostBook(c echo.Context) error {
	request := dto.BookRequest{}
	if err := c.Bind(&request); err != nil {
		return errResp(c, 400, "Invalid request", err)
	}
	if err := c.Validate(&request); err != nil {
		return errResp(c, 400, "Invalid request", err)
	}

	book, err := h.svc.CreateBook(request)
	if err != nil {
		return errResp(c, 500, "Server error", err)
	}

	return c.JSON(201, book)
}

func (h *Handler) GetAllBooks(c echo.Context) error {
	books, err := h.svc.GetBooks()
	if err != nil {
		return errResp(c, 500, "Server error", err)
	}

	return c.JSON(200, books)
}

func (h *Handler) GetSingleBook(c echo.Context) error {
	id := c.Param("id")
	book, err := h.svc.GetBookById(id)
	if err != nil {
		return errResp(c, 404, "Book not found", err)
	}

	return c.JSON(200, book)
}

func (h *Handler) UpdateSingleBook(c echo.Context) error {
	request := dto.BookRequest{}
	if err := c.Bind(&request); err != nil {
		return errResp(c, 400, "Invalid request", err)
	}
	if err := c.Validate(&request); err != nil {
		return errResp(c, 400, "Invalid request", err)
	}

	id := c.Param("id")
	book, err := h.svc.UpdateBookById(id, request)
	if err != nil {
		return errResp(c, 404, "Book not found", err)
	}

	return c.JSON(200, book)
}

func (h *Handler) DeleteSingleBook(c echo.Context) error {
	id := c.Param("id")
	err := h.svc.DeleteBookById(id)
	if err != nil {
		return errResp(c, 404, "Book not found", err)
	}

	return c.String(204, "")
}
