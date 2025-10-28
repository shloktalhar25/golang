
// Gin + GORM CRUD API Example ,store data in SQLite database , test with Postman.

// make Book API real-world ready by connecting it to a database 
// using GORM.

// In this projet we shall learn to Connect to SQLite using GORM

//GORM (Go Object Relational Mapping) is a popular ORM library for Go that simplifies database interactions.

// learnings:

// How to connect your Gin app to a database
// How to use GORM (Go ORM) to manage data
// How to perform CRUD operations directly in the DB (instead of in-memory slices)


package main 

import (
	"net/http"
	"github.com/gin-gonic/gin" //for gin framework
	"gorm.io/driver/sqlite" // for sqlite driver
	"gorm.io/gorm" // for gorm (go Object Relational Mapping)
)

// Book struct - > represent our data model
type Book struct {
	ID     int  `json:"id" gorm:"primaryKey"` //used by gorm to identify primary key and auto increment , you dont need to set it manually the database will handle it
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  int `json:"price"`
}

// Database setup 

var DB *gorm.DB // global variable to hold db connection so that we can use it in handler functions



// ConnectDatabase connects GORM to SQLite and auto-migrates tables
func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&Book{})
}



// Handler functions

// GET /books -> get all the books
func getBooks(c *gin.Context) {
	var books []Book // slice to hold books from db
	DB.Find(&books) // fetch all books from database
	c.IndentedJSON(http.StatusOK,books) // return books as json response
	//the flow here is : first we define a slice to hold the books fetched from db
	// then we use gorm's Find method to get all records from books table and store them in the slice
	// finally we return the slice as json response with 200 status code
}

// Handler to add a new book to the database

func addBook(c *gin.Context) {
	var newBook Book
	//bind the json payload to newBook struct and check for errors
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// add the new book to the database
	DB.Create(&newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
	//flow here is : first we bind the incoming json payload to newBook struct
	// then we use gorm's Create method to insert the new book record into the books table
	// finally we return the newly created book as json response with 201 status code ,creaation of new resource is 201 status code
}






// Handler to update an existing book in the database\

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book // struct to hold the existing book

	// first find the book by id and check if it exists , if not return 404, use First method of gorm to get first record that matches the condition
	if err := DB.First(&book,id).Error ; err != nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}


	// create a struct to hold the updated data
	var updatedData Book

	// bind the json payload to updatedData struct and check for errors
	// simple words : take the data from the request body and put it in updatedData struct
	if err := c.BindJSON(&updatedData);err != nil {
		return
	}

	book.Title = updatedData.Title
	book.Author = updatedData.Author
	book.Price = updatedData.Price

	DB.Save(&book) // save the updated book back to the database

	c.IndentedJSON(http.StatusOK,book) // return the updated book as json response
	//flow here is : first we get the book by id using gorm's First method
	// then we bind the incoming json payload to updatedData struct
	// we update the fields of the existing book with the new data
	// finally we save the updated book back to the database using gorm's Save method and return it as json response with 200 status code

}

// Handler to delete a book from the database
func deleteBook(c *gin.Context) {
	id :=c.Param("id")

	var book Book 

	// first find the book by id and check if it exists , if not return 404
	// we have to find the book first before deleting it
	if err := DB.First(&book,id).Error ; err != nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found"})
		return
	}

	//Delete the book from the database
	DB.Delete(&book)
	c.IndentedJSON(http.StatusOK,gin.H{"message":"Book deleted successfully"}) // indentedJSON is used to return json response as pretty printed json
}





// Main function
func main() {
	router := gin.Default() // create a gin router with default middleware

	ConnectDatabase() // connect to the database

	//define routes and their handler functions
	router.GET("/books",getBooks)// get all books
	router.POST("/books",addBook) // add a new book
	router.PUT("/books/:id",updateBook) // update a book by id
	router.DELETE("/books/:id",deleteBook) // delete a book by id

	router.Run() // start the server on default port 8080
}