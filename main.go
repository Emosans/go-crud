package main

import (
	"tr/models"

	"github.com/gin-gonic/gin"
)

func main(){
	// l:=models.Library{}

	// l.Booktype(models.Individual{Person: "Alice",Type: models.BORROW,Time: 1})
	// l.UpdateTime("Alice",models.Individual{Person: "Alice",Type: models.BORROW,Time: 2},2)
	// l.UpdatePerson(models.Individual{Person:"John",Type: models.BORROW,Time: 2},"Alice")

	router := gin.Default()

	router.GET("/books",models.GetBooks)
	router.POST("/add",models.AddBook)
	router.POST("/update/:person",models.ReturnBook)
	router.DELETE("/delete/:person",models.DeleteBook)

	router.Run("localhost:8080")
}