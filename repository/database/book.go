package database

import (
	"project_structure/config"
	"project_structure/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetBooks() (books []model.Book, err error)
	GetBook(id uint) (book model.Book, err error)
	UpdateBook(book *model.Book) error
	DeleteBook(book *model.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) CreateBook(book *model.Book) error {
	if err := config.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) GetBooks() (books []model.Book, err error) {
	if err = config.DB.Find(&books).Error; err != nil {
		return
	}
	return
}

func (b *bookRepository) GetBook(id uint) (book model.Book, err error) {
	book.ID = id
	if err = config.DB.First(&book).Error; err != nil {
		return
	}
	return
}

func (b *bookRepository) UpdateBook(book *model.Book) error {
	if err := config.DB.Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) DeleteBook(book *model.Book) error {
	if err := config.DB.Delete(book).Error; err != nil {
		return err
	}
	return nil
}
