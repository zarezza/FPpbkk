package models

import (
	"database/sql"
	"final-project/config"
	"final-project/entities"
	"fmt"
)

type BookModel struct {
	conn *sql.DB
}

func NewBookModel() *BookModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &BookModel{
		conn: conn,
	}
}

func (b *BookModel) FindAll() ([]entities.Book, error) {
	rows, err := b.conn.Query("select * from books")
	if err != nil {
		return []entities.Book{}, err
	}
	defer rows.Close()

	var dataBooks []entities.Book
	for rows.Next() {
		var book entities.Book
		rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Publisher,
			&book.ISBN,
			&book.Year,
			&book.Category)

		dataBooks = append(dataBooks, book)
	}
	return dataBooks, nil
}

func (b *BookModel) Create(book entities.Book) bool {
	result, err := b.conn.Exec("insert into books (title, author, publisher, isbn, year, category) values(?,?,?,?,?,?)",
		book.Title, book.Author, book.Publisher, book.ISBN, book.Year, book.Category)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (b *BookModel) Find(id int64, book *entities.Book) error {
	return b.conn.QueryRow("select * from books where id = ?", id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.ISBN,
		&book.Year,
		&book.Category,
	)
}

func (b *BookModel) Update(book entities.Book) error {
	_, err := b.conn.Exec(
		"update books set title=?, author=?, publisher=?, isbn=?, year=?, category=? where id=?",
		book.Title, book.Author, book.Publisher, book.ISBN, book.Year, book.Category, book.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (b *BookModel) Delete(id int64) {
	b.conn.Exec("delete from books where id = ?", id)
}
