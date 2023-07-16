package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/username/myapp/app/model"
)

func GetAllBooks(c echo.Context) error {
	books, err := model.GetAllBooks()
	if err != nil {
		response := APIResponse{
			Kode:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to retrieve books",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    books,
	}

	return c.JSON(http.StatusOK, response)
}



func GetBookByID(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	book, err := model.GetBookByID(bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			response := APIResponse{
				Kode:    http.StatusNotFound,
				Status:  "error",
				Message: "Book not found",
				Data:    nil,
			}
			return c.JSON(http.StatusNotFound, response)
		}
		response := APIResponse{
			Kode:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to retrieve book",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    book,
	}

	return c.JSON(http.StatusOK, response)
}


func CreateBook(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		response := APIResponse{
			Kode:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	bookID, err := model.CreateBook(book)
	if err != nil {
		response := APIResponse{
			Kode:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to create book",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := APIResponse{
		Kode:    http.StatusCreated,
		Status:  "success",
		Message: "Book created successfully",
		Data:    bookID,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	book, err := model.GetBookByID(bookID)
	if err != nil {
		response := APIResponse{
			Kode:    http.StatusNotFound,
			Status:  "error",
			Message: "Book not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	updatedBook := new(model.Book)
	if err := c.Bind(updatedBook); err != nil {
		response := APIResponse{
			Kode:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author

	err = model.UpdateBook(&book)
	if err != nil {
		response := APIResponse{
			Kode:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to update book",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Book updated successfully",
		Data:    book,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	err := model.DeleteBook(bookID)
	if err != nil {
		response := APIResponse{
			Kode:    http.StatusInternalServerError,
			Status:  "error",
			Message: "Failed to delete book",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Book deleted successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, response)
}


func parseID(id string) int {
	bookID, _ := strconv.Atoi(id)
	return bookID
}