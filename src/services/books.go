package services

import (
	"desent/src/dto"
	"desent/src/entity"
	"desent/src/utils"
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

func (s *Service) GetBooks(q dto.BookQueryParam) ([]entity.Book, error) {
	db := s.app.DB
	books := []entity.Book{}

	prep := db.Model(&entity.Book{})
	if len(q.Author) > 0 {
		prep = prep.Where("author LIKE ?", utils.Like(q.Author))
	}

	if q.Limit > 0 {
		if q.Page <= 0 {
			q.Page = 1
		}

		offset := (q.Page - 1) * q.Limit
		prep = prep.Offset(offset).Limit(q.Limit)
	}

	if err := prep.Find(&books).Error; err != nil {
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
