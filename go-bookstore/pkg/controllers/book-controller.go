package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abhishek-1705/go-bookstore/pkg/models"
	"github.com/Abhishek-1705/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book //creating a variable of type Book Struture for the operations

//controller Function to get Book  Data
func GetBook(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBook()                    //Getting all the Book Data from the model Function, return a Slice of Book Data
	res, _ := json.Marshal(newBooks)                   //Converting it to the json format, since it is a slice of Data
	w.Header().Set("Content-Type", "pkglication/json") //setting the Header Content type to json
	w.WriteHeader(http.StatusOK)                       //writting the header Status, 200 is StatusOK
	w.Write(res)                                       //wrting the response with the Book Data

}

//controller Function to get book Data by Id
func GetBookById(w http.ResponseWriter, r *http.Request) {

	Vars := mux.Vars(r)                       //storing the varibles inside the request to Vars
	bookId := Vars["bookId"]                  //Storing the request Vars BookId
	ID, err := strconv.ParseInt(bookId, 0, 0) //passing the bookid, checking any other edge case
	if err != nil {
		fmt.Println("error while passing") //logging out the error
	}

	bookDetails, _ := models.GetBookById(ID)           //calling the GETBookbyid function for the models
	res, _ := json.Marshal(bookDetails)                //converting the Book data Struture to json format
	w.Header().Set("content-Type", "pkglication/json") //setting the Header Content type to json
	w.WriteHeader(http.StatusOK)                       //writting the header Status, 200 is StatusOK
	w.Write(res)                                       //wrting the response with the Book Data

}

//controller Function to Create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}                       //creating a variable for the reference of the Book Data Structure form models
	utils.ParseBody(r, CreateBook)                     //passing the request and the Book Data Structure to ParseBody function from the utils package
	b := CreateBook.CreateBook()                       //calling the CreateBook function from the models package
	res, _ := json.Marshal(b)                          //converting the Book data Struture to json format
	w.Header().Set("content-Type", "pkglication/json") //setting the Header Content type to json
	w.WriteHeader(http.StatusOK)                       //writting the header Status, 200 is StatusOK
	w.Write(res)                                       //wrting the response with the Book Data
}

//controller Funtion for the Deletion of Book data
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       //storing the varibles inside the request to Vars
	bookId := vars["bookId"]                  //Storing the request Vars BookId
	ID, err := strconv.ParseInt(bookId, 0, 0) //passing the bookid, checking any other edge case
	if err != nil {
		fmt.Println("error While Parsing") //logging out the error
	}
	book := models.DeleteBook(ID)                      //calling the Delete Function form the models
	res, _ := json.Marshal(book)                       //converting the responce to the Json
	w.Header().Set("Content-type", "pkglication/json") //setting the Header Content type to json
	w.Write(res)                                       //wrting the response with the Book Data

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var UpdateBook = &models.Book{} //Making a Temp Variable for the updatation
	utils.ParseBody(r, UpdateBook)  //parsing the respone to Updatebook models

	vars := mux.Vars(r)                       //storing the varibles inside the request to Vars
	bookId := vars["bookId"]                  //Storing the request Vars BookId
	ID, err := strconv.ParseInt(bookId, 0, 0) //passing the bookid, checking any other edge case
	if err != nil {
		fmt.Println("errro while parsing") //logging out the error
	}

	bookDetails, db := models.GetBookById(ID) //getting the book details for the databases, by calling the GETBookbyid form models
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name //updating the book name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author //updating the book author name
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication //updating the book Pulication data
	}

	db.Save(&bookDetails) //saving the updated  Book data in database

	res, _ := json.Marshal(bookDetails)                //converting the responce to the Json
	w.Header().Set("Content-Type", "pkglication/json") //setting the Header Content type to json
	w.WriteHeader(http.StatusOK)                       //writting the header Status, 200 is StatusOK
	w.Write(res)                                       //wrting the response with the Book Data
}
