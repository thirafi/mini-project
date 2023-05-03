package controller

import (
	"net/http"
	"project_structure/model"
	"project_structure/repository/database"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController interface {
	GetBookcontroller(c echo.Context) error
	GetBookController(c echo.Context) error
	CreateBookController(c echo.Context) error
	DeleteBookController(c echo.Context) error
	UpdateBookController(c echo.Context) error
}

type bookController struct {
	bookUsecase    usecase.BookUsecase
	bookRepository database.BookRepository
}

func NewBookController(
	bookUsecase usecase.BookUsecase,
	bookRepository database.BookRepository,
) *bookController {
	return &bookController{
		bookUsecase,
		bookRepository,
	}
}

func (b *bookController) GetBookcontroller(c echo.Context) error {
	books, e := b.bookRepository.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func (b *bookController) GetBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := b.bookRepository.GetBook(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

// create new book
func (b *bookController) CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := b.bookUsecase.CreateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func (b *bookController) DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.bookUsecase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete book",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf buku tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func (b *bookController) UpdateBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book := model.Book{}
	c.Bind(&book)
	book.ID = uint(id)
	if err := b.bookUsecase.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update book",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf buku tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
