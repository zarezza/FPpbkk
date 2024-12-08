package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateAndSeed(db *sql.DB) error {
	err := migrate(db)
	if err != nil {
		return err
	}

	err = seed(db)
	if err != nil {
		return err
	}

	return nil
}

func migrate(db *sql.DB) error {
	// Drop the books table if it exists
	_, err := db.Exec("DROP TABLE IF EXISTS books")
	if err != nil {
		return err
	}
	fmt.Println("Dropped table if it existed.")

	// Create the books table
	createTableQuery := `
    CREATE TABLE books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        published_date DATE,
        isbn VARCHAR(13),
        pages INT,
        cover VARCHAR(255),
        language VARCHAR(50)
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	fmt.Println("Created table books.")
	return nil
}

func seed(db *sql.DB) error {
	// Insert sample data into the books table
	seedQuery := `
    INSERT INTO books (title, author, published_date, isbn, pages, cover, language) VALUES
    ('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', '9780743273565', 180, 'great_gatsby.jpg', 'English'),
    ('To Kill a Mockingbird', 'Harper Lee', '1960-07-11', '9780061120084', 281, 'to_kill_a_mockingbird.jpg', 'English'),
    ('1984', 'George Orwell', '1949-06-08', '9780451524935', 328, '1984.jpg', 'English');`

	_, err := db.Exec(seedQuery)
	if err != nil {
		return err
	}
	fmt.Println("Seeded table books with sample data.")
	return nil
}
