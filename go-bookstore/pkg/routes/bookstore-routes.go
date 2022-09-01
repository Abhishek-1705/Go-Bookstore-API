package routes //Importing the package

import (
	"github.com/Abhishek-1705/go-bookstore/pkg/controllers" //Importing the the controller package through a github control
	"github.com/gorilla/mux"                                //Importing the Mux libary for Routing
)

//Creating a Function for the Request Handler to call the function from the controller package using a router as Mux Router.
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           //Handler to Create Book in Database
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               //Handler for to Book Data
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   //Handler to get book by Book Id
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    //Handler for the updating the Book Data
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") //Handler for deleting the book Data

}
