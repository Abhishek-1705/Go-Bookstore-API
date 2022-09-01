//Package for connecting the Data Base

package config

import (
	"github.com/jinzhu/gorm"                  //Gorm package for Connecting the Data Base
	_ "github.com/jinzhu/gorm/dialects/mysql" //Importing the Mysql libary from Gorm Package for the connection.
)

var (
	db *gorm.DB //Creating a pointer for the Database of type Gorm.bd
)

func Connect() {
	d, err := gorm.Open("mysql", "UserName:password/simplerest?charset=utf8&parseTime=True&loc=Local") //Opening the Connection for Data Base, ("DataBase type","Owner name : password /table name?Perdefine Char for mysql Database and Json")
	if err != nil {                                                                                    // catching the error
		panic(err) //Logging out the Error
	}
	db = d //If Connection made, then Pointer is equal to databse connection

}

func GetDB() *gorm.DB { //setter function for the DataBase and Pointer
	return db
}
