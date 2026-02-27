package services

import (
	"desent/src/dto"
	"desent/src/entity"
)

func (s *Service) CreateBook(req dto.BookRequest) (*entity.Book, error) {
	db := s.app.DB
	bookData := entity.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := db.Create(&bookData).Error; err != nil {
		return nil, err
	}

	return &bookData, nil
}

func (s *Service) GetBooks() ([]entity.Book, error) {
	db := s.app.DB
	books := []entity.Book{}

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) GetBookById(bookId string) (*entity.Book, error) {
	db := s.app.DB
	book := entity.Book{}

	if err := db.Where("id = ?", bookId).Take(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *Service) UpdateBookById(id string, request dto.BookRequest) (*entity.Book, error) {
	db := s.app.DB

	book, err := s.GetBookById(id)
	if err != nil {
		return nil, err
	}

	book.Title = request.Title
	book.Year = request.Year
	book.Author = request.Author
	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) DeleteBookById(id string) error {
	db := s.app.DB

	book, err := s.GetBookById(id)
	if err != nil {
		return err
	}

	if err := db.Unscoped().Model(&entity.Book{}).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
