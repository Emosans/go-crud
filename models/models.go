package models

import (
	// "fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type BookType string

const (
	BORROW BookType = "Borrow"
	RETURN BookType = "Returned"
)

// each person trade info
type Individual struct {
	Person string
	Type   BookType
	Time   int
}

// portfolio info
type Library struct {
	Libraries map[string][]Individual
	Fine      map[string][]int
}

// func (l *Library) Booktype(p Individual) {
// 	// check type borrow (if yes add the person)
// 	// add the person to the library libraries array type(Individual)
// 	// specify time for borrow

// 	if l.Libraries == nil {
// 		l.Libraries = make(map[string][]Individual)
// 	}

// 	// add a person
// 	l.Libraries[p.Person] = append(l.Libraries[p.Person], p)
// 	fmt.Println(l.Libraries)
// }

// func (l *Library) UpdateTime(person string, p Individual, time int) {
// 	if libraries, ok := l.Libraries[person]; ok {
// 		for i, library := range libraries {
// 			if library.Person == p.Person {
// 				l.Libraries[person][i] = p
// 			}
// 		}
// 	}
// 	fmt.Println(l.Libraries[person])
// }

// func (l *Library) UpdatePerson(p Individual,person string){

// 	if(p.Person==person){
// 		if libraries,ok := l.Libraries[person]; ok{
// 			for i,library := range libraries {
// 				if library.Person == person {
// 					l.Libraries[person][i] = p
// 				}
// 			}
// 		}
// 	} else {
// 		if libraries,ok := l.Libraries[person]; ok{
// 			for i,library := range libraries {
// 				l.Libraries[person] = append(libraries[:i],libraries[i+1:]... )
// 				fmt.Println(library)
// 				if l.Libraries[p.Person] == nil {
// 					l.Libraries[p.Person] = []Individual{}
// 				}
// 				l.Libraries[p.Person] = append(l.Libraries[p.Person], p)
// 			}
// 		}
// 	}
// 	fmt.Println(l.Libraries)
// }

// apis

var books []Individual

func GetBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func AddBook(c *gin.Context){
	var newBook Individual
	if err:=c.BindJSON(&newBook); err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Cant add book"})
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func ReturnBook(c *gin.Context){
	var person = c.Param("person")
	// err:=c.BindJSON(&bookinfo)

	for i,book := range books{
		if book.Person == person{
			books[i].Type = "Returned"
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"msg":"updated return record"})
}

func DeleteBook(c *gin.Context){
	var person = c.Param("person")

	for i,book := range books{
		if book.Person == person{
			books = append(books[:i],books[i+1:]... )
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"msg":"deleted"})
}

func GetPeople(c *gin.Context){
	var person []string

	for _,book := range books{
		person = append(person, book.Person)
	}

	c.IndentedJSON(http.StatusOK,person)

}

func UpdateTime(c *gin.Context){
	var param = c.Param("time")
	time,err := strconv.Atoi(param)

	if err!=nil{
		return
	}
	var updatePerson Individual
	if err := c.BindJSON(&updatePerson); err!=nil {
		return
	}

	for i,book := range books{
		if book.Person == updatePerson.Person{
			books[i].Time = time
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"msg":"updated time"})

}