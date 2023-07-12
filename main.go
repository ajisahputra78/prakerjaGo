package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = map[int]Book{
	1: Book{ID: 1, Title: "Book 1", Author: "Author 1"},
	2: Book{ID: 2, Title: "Book 2", Author: "Author 2"},
}

func getAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}