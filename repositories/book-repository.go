package repositories

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/guilhermebehs/book-crud-in-go/entities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type BookRepository struct {
	db *sql.DB
}

func CreateRepository() *BookRepository {
	bookRepository := BookRepository{}
	bookRepository.Init()
	return &bookRepository
}

func connectDb() *sql.DB {
	dotenvErr := godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	if dotenvErr != nil {
		panic("Error loading config file: " + dotenvErr.Error())
	}
	db, errConn := sql.Open("postgres", "postgresql://"+user+":"+pass+"@"+host+":"+port+"/"+dbName+"?sslmode=disable")
	if errConn != nil {
		panic("Error connecting db: " + errConn.Error())
	}
	return db
}

func (br *BookRepository) Init() {

	br.db = connectDb()

	_, err := br.db.Query(`
	CREATE TABLE IF NOT EXISTS book
	  (
	   id SERIAL PRIMARY KEY, 
	   isbn VARCHAR UNIQUE NOT NULL, 
	   title VARCHAR NOT NULL,
	   year VARCHAR NOT NULL,
	   author VARCHAR NOT NULL
	   )`)
	if err != nil {
		log.Fatal("Error creating table: " + err.Error())
	}
}

func (br *BookRepository) GetByIsbn(isbn string) (entities.Book, error) {
	rows, err := br.db.Query("SELECT isbn, title, year, author FROM book WHERE isbn = $1", isbn)
	if err != nil {
		return entities.Book{}, err
	}

	book := entities.Book{}
	for rows.Next() {
		rows.Scan(&book.Isbn, &book.Title, &book.Year, &book.Author)
		break
	}
	if book.Isbn == "" {
		return book, errors.New("not found")
	} else {
		return book, nil
	}
}

func (br *BookRepository) Create(book entities.Book) error {
	_, err := br.db.Query("INSERT INTO book (isbn, title, year, author) VALUES($1,$2,$3,$4)", book.Isbn, book.Title, book.Year, book.Author)
	if err != nil {
		return err
	} else {
		return nil
	}

}
