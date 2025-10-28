// We’ll build a Book Management API step-by-step.
// You’ll learn routing, JSON handling, and how to interact with data (in-memory first,
// then we can move to a database later).
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book struct - > represrnt our data model

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// In memory book stoarage

var books = []Book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 10.99},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 8.99},
	{ID: "3", Title: "1984", Author: "George Orwell", Price: 9.99},
}

// as of now we have defined our data model and in-memory storage
// Next, we will set up the Gin router and define endpoints for CRUD operations on books.

// GET /books -> get all the books

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // return all books as json response(response )
}

// GET /books/:id -> get book by id
func getBooksByID(c *gin.Context) {
	id := c.Param("id")
	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b) // if found return book as json response with 200 status
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) //if not found return 404 status with message
	// we use gin.H to simply create a map for json response

}

// POST /books -> add a new book
func addBook(c *gin.Context) {
	var newBook Book
	//bind the json payload to newBook struct
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// else add the new book to the in-memory storage
	books = append(books, newBook)
	// return the newly added book as json response with 201 status
	c.IndentedJSON(http.StatusCreated, newBook)
}

// PUT /books/:id -> update book by id
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	// bind the json payload to updatedBook struct
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	// else find the book by id and update it
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})

}
func deleteBook(c *gin.Context) {
	// get the book id from url param
	id := c.Param("id")
	// find the book by id and delete it
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// main function to setup router and start server
func main() {

	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBooksByID)
	router.POST("/books", addBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	// router run on default port 8080
	router.Run("0.0.0.0:8080")

}
