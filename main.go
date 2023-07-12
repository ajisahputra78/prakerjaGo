package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = map[int]Book{
	1: Book{ID: 1, Title: "Buku 1", Author: "Penulis 1"},
	2: Book{ID: 2, Title: "Buku 2", Author: "Penulis 2"},
}

type APIResponse struct {
	Kode	int16 `json:"kode"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func getAllBooks(c echo.Context) error {
	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    books,
	}

	return c.JSON(http.StatusOK, response)
}

func getBookByID(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	book, exists := books[bookID]
	if !exists {
		response := APIResponse{
			Kode:    http.StatusNotFound,
			Status:  "error",
			Message: "Book not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    book,
	}

	return c.JSON(http.StatusOK, response)
}

func createBook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		response := APIResponse{
			Kode:    http.StatusBadRequest,
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	book.ID = len(books) + 1
	books[book.ID] = *book

	response := APIResponse{
		Kode:    http.StatusCreated,
		Status:  "success",
		Message: "Book created successfully",
		Data:    book,
	}

	return c.JSON(http.StatusCreated, response)
}

func updateBook(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	book, exists := books[bookID]
	if !exists {
		response := APIResponse{
			Kode:    http.StatusNotFound,
			Status:  "error",
			Message: "Book not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	updatedBook := new(Book)
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

	books[bookID] = book

	response := APIResponse{
		Kode:    http.StatusOK,
		Status:  "success",
		Message: "Book updated successfully",
		Data:    book,
	}

	return c.JSON(http.StatusOK, response)
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	bookID := parseID(id)

	_, exists := books[bookID]
	if !exists {
		response := APIResponse{
			Kode:    http.StatusNotFound,
			Status:  "error",
			Message: "Book not found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	delete(books, bookID)

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

func main() {
	e := echo.New()

	// Routes
	e.GET("/api/books", getAllBooks)
	e.GET("/api/books/:id", getBookByID)
	e.POST("/api/books/create", createBook)
	e.PUT("/api/books/:id/update", updateBook)
	e.DELETE("/api/books/:id/delete", deleteBook)

	// Start server
	e.Start(":8000")
}