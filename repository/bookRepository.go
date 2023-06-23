package repository

import (
	"learn-clean-arch/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	FindByID(bookID int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(book models.Book) (models.Book, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db}
}

func (r *BookRepositoryImpl) FindAll() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *BookRepositoryImpl) FindByID(bookID int) (models.Book, error) {
	var book models.Book

	err := r.db.Find(&book, bookID).Error

	return book, err
}

func (r *BookRepositoryImpl) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *BookRepositoryImpl) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *BookRepositoryImpl) Delete(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
