package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type BookUsecase interface {
	CreateBook(book *model.Book) error
	GetBook(id uint) (book model.Book, err error)
	GetListBooks() (books []model.Book, err error)
	UpdateBook(book *model.Book) (err error)
	DeleteBook(id uint) (err error)
}

type bookUsecase struct {
	bookRepository database.BookRepository
}

func NewBookUsecase(bookRepo database.BookRepository) *bookUsecase {
	return &bookUsecase{bookRepository: bookRepo}
}

func (b *bookUsecase) CreateBook(book *model.Book) error {

	// check title cannot be empty
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}

	// check creator
	if book.Creator == "" {
		return errors.New("book creator cannot be empty")
	}

	err := b.bookRepository.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookUsecase) GetBook(id uint) (book model.Book, err error) {
	book, err = b.bookRepository.GetBook(id)
	if err != nil {
		fmt.Println("GetBook: Error getting book from database")
		return
	}
	return
}

func (b *bookUsecase) GetListBooks() (books []model.Book, err error) {
	books, err = b.bookRepository.GetBooks()
	if err != nil {
		fmt.Println("GetListBooks: Error getting books from database")
		return
	}
	return
}

func (b *bookUsecase) UpdateBook(book *model.Book) (err error) {
	err = b.bookRepository.UpdateBook(book)
	if err != nil {
		fmt.Println("UpdateBook : Error updating book, err: ", err)
		return
	}

	return
}

func (b *bookUsecase) DeleteBook(id uint) (err error) {
	book := model.Book{}
	book.ID = id
	err = b.bookRepository.DeleteBook(&book)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book, err: ", err)
		return
	}

	return
}
