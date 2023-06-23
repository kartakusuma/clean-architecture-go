package services

import (
	"learn-clean-arch/models"
	"learn-clean-arch/repository"
)

type BookService interface {
	FindAll() ([]models.Book, error)
	FindByID(bookID int) (models.Book, error)
	Create(bookInput models.BookInput) (models.Book, error)
	Update(bookInput models.BookInput, bookID int) (models.Book, error)
	Delete(bookID int) (models.Book, error)
}

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookrepository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookrepository}
}

func (s *BookServiceImpl) FindAll() ([]models.Book, error) {
	books, err := s.bookRepository.FindAll()
	return books, err
}

func (s *BookServiceImpl) FindByID(bookID int) (models.Book, error) {
	book, err := s.bookRepository.FindByID(bookID)
	return book, err
}

func (s *BookServiceImpl) Create(bookInput models.BookInput) (models.Book, error) {
	newBook := models.Book{
		Title:  bookInput.Title,
		Author: bookInput.Author,
	}

	createdBook, err := s.bookRepository.Create(newBook)
	return createdBook, err
}

func (s *BookServiceImpl) Update(bookInput models.BookInput, bookID int) (models.Book, error) {
	book, err := s.bookRepository.FindByID(bookID)
	if err != nil {
		return book, err
	}

	book.Title = bookInput.Title
	book.Author = bookInput.Author

	updatedBook, err := s.bookRepository.Update(book)

	return updatedBook, err
}

func (s *BookServiceImpl) Delete(bookID int) (models.Book, error) {
	book, err := s.bookRepository.FindByID(bookID)
	if err != nil {
		return book, err
	}

	deletedBook, err := s.bookRepository.Delete(book)

	return deletedBook, err
}
