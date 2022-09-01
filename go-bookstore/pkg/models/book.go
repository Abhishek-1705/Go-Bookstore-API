package models

//Importing the libaries
import (
	"github.com/Abhishek-1705/go-bookstore/pkg/config" //The Config file that returnes the Data base Connection
	"github.com/jinzhu/gorm"                           //gorm for the Data base Connection
)

//Variable for storting the Data base connection
var db *gorm.DB

//Defining the Sturture of The Book for the opertaion
type Book struct {
	gorm.Model         //Adding the Model from the Gorm Libary
	Name        string `gorm:"json:name"`   //name of the Book
	Author      string `json:"author"`      //Author of the book
	Publication string `json:"publication"` //Publication of the book
}

//initialiation Function for the database
func init() {
	config.Connect()        //calling the connection function for the Config File
	db = config.GetDB()     //getting the Database connection
	db.AutoMigrate(&Book{}) //Making the a table for the Book
}

////Function for the Data Based Operations

//function for Creating the Book
//Gorm has all the function for the Data base Operation,it use SQl Query in the Backend
func (b *Book) CreateBook() *Book { //passing the Book Struture and a pointer to new book recored entered
	db.NewRecord(b) //making the new Book Record  using the Gorm NewRecord Function passing the Book Data to Datbase
	db.Create(&b)   // Creating the  New Record in the Data Base
	return b        //Returning the New Book Data pointer
}

//Function for Getting Book Data
func GetAllBook() []Book { //Return the book Data as Slice of Struture
	var Books []Book //making a Temp Variable for the Book Data
	db.Find(&Books)  //using the GORM Find Function to get book data
	return Books
}

//Function for getting Book Data by Id
func GetBookById(Id int64) (*Book, *gorm.DB) { //passing the Id of the Book and Returing the Book as Struture and a db connection
	var getbook Book                          //Temp Variable for the Returning of the Function
	db := db.Where("Id=?", Id).Find(&getbook) //finding the book with book id, Where function return the interface passed value with argument passed value
	return &getbook, db                       //returning the book Data and db pointer
}

// Function for Deleting the book for the Databased
func DeleteBook(Id int64) Book {
	var book Book                     //temp varabile for returing the Functions
	db.Where("Id=?", Id).Delete(book) //Deleting the book with book id, Where function return the interface passed value with argument passed value
	return book                       //returing the Deleted Book Data

}
