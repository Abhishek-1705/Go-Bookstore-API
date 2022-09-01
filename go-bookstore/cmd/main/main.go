//Defining the all the Server and Local Host
//Defining all the other package for the process

package main

//Importing all tha libaries and modules
import (
	"log"      //for logging out the data
	"net/http" //for the Http routes and Handler

	"github.com/Abhishek-1705/go-bookstore/pkg/routes" //importing the Routes package
	"github.com/gorilla/mux"                           //For Routing a the Functions
	_ "github.com/jinzhu/gorm/dialects/mysql"          //for the MySQL
)

//The main Functions
func main() {
	r := mux.NewRouter()              //Creating a New Router for Mux
	routes.RegisterBookStoreRoutes(r) //Calling the Routes  Handler Function for Futher Routing

	http.Handle("/", r) //Passing the Request to Handler Functions

	log.Fatal(http.ListenAndServe(":5050", r)) //Stating the Local Host server at port 5050

}
