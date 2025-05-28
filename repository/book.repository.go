package repository

import (
	"github.com/Heinirich/gorm-mysql/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.BookModel) error
	FindAll() ([]models.BookModel, error)
	FindByID(id uint) (models.BookModel, error)
	Update(book *models.BookModel) error
	Delete(id uint) error
}

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepo{db}
}

func (r *bookRepo) Create(book *models.BookModel) error {
	return r.db.Create(book).Error
}

func (r *bookRepo) FindAll() ([]models.BookModel, error) {
	var books []models.BookModel
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepo) FindByID(id uint) (models.BookModel, error) {
	var book models.BookModel
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepo) Update(book *models.BookModel) error {
	return r.db.Save(book).Error
}

func (r *bookRepo) Delete(id uint) error {
	return r.db.Delete(&models.BookModel{}, id).Error
}
