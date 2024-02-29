package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string  `jason:id`
	Title    string  `jason:title`
	Author   string  `jason:author`
	Quantity int     `jason:quantity`
}

//Books slice for api
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found.")
}

func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}