package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

// to store book data in memory
//json:"id" = specify what a field name should be when struct content are serialized into JSON
type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

// data in memory. typical api interact with database 
// book slice to seed record book data 
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// get request, handler function - responds with the list of all books as JSON
// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
//In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context){
	var newBook book 

	// call BindJSON to bind the received JSON to newBook
	if err:= c.BindJSON(&newBook); err!=nil{
		return
	}
	// add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	// gin router setup
	router:= gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.Run("localhost:8080")

}