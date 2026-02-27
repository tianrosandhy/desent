package handlers

import (
	"desent/src/dto"
	"desent/src/pkg/response"
	"log"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PostBook(c echo.Context) error {
	request := dto.BookRequest{}
	if err := c.Bind(&request); err != nil {
		return response.ErrorResponse("Invalid request", err).Send(c, 400)
	}
	if err := c.Validate(&request); err != nil {
		return response.ErrorResponse("Invalid request", err).Send(c, 400)
	}

	book, err := h.svc.CreateBook(request)
	if err != nil {
		log.Printf("ERR CreateBook = %+v", err)
		return response.ErrorResponse("Server error").Send(c, 500)
	}

	return c.JSON(201, book)
}

func (h *Handler) GetAllBooks(c echo.Context) error {
	q := dto.BookQueryParam{}
	c.Bind(&q)

	books, err := h.svc.GetBooks(q)
	if err != nil {
		log.Printf("ERR CreateBooks = %+v", err)
		return response.ErrorResponse("Server error").Send(c, 500)
	}

	return c.JSON(200, books)
}

func (h *Handler) GetSingleBook(c echo.Context) error {
	id := c.Param("id")
	book, err := h.svc.GetBookById(id)
	if err != nil {
		return response.ErrorResponse("Book not found").Send(c, 404)
	}

	return c.JSON(200, book)
}

func (h *Handler) UpdateSingleBook(c echo.Context) error {
	request := dto.BookRequest{}
	if err := c.Bind(&request); err != nil {
		return response.ErrorResponse("Invalid request", err).Send(c, 400)
	}
	if err := c.Validate(&request); err != nil {
		return response.ErrorResponse("Invalid request", err).Send(c, 400)
	}

	id := c.Param("id")
	book, err := h.svc.UpdateBookById(id, request)
	if err != nil {
		return response.ErrorResponse("Book not found").Send(c, 404)
	}

	return c.JSON(200, book)
}

func (h *Handler) DeleteSingleBook(c echo.Context) error {
	id := c.Param("id")
	err := h.svc.DeleteBookById(id)
	if err != nil {
		return response.ErrorResponse("Book not found").Send(c, 404)
	}

	return c.String(204, "")
}
