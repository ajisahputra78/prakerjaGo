package model

import (
	"database/sql"
	"log"
)

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

func InitBookModel(database *sql.DB) {
	db = database
}

func GetAllBooks() ([]Book, error) {
	query := "SELECT * FROM books"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}

	if len(books) == 0 {
		if err == sql.ErrNoRows {
			return []Book{}, nil
		} else {
			log.Println(err)
			return nil, err
		}
	}

	return books, nil
}

func GetBookByID(id int) (Book, error) {
	query := "SELECT * FROM books WHERE id = ?"

	row := db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		log.Println(err)
		return book, err
	}

	return book, nil
}

func CreateBook(book *Book) (int64, error) {
	query := "INSERT INTO books (title, author) VALUES (?, ?)"

	result, err := db.Exec(query, book.Title, book.Author)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	bookID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return bookID, nil
}

func UpdateBook(updatedBook *Book) error {
	query := "UPDATE books SET title = ?, author = ? WHERE id = ?"

	_, err := db.Exec(query, updatedBook.Title, updatedBook.Author, updatedBook.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
