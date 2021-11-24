package database

import "fmt"

var connection string

// automatic running in first time called package
func init() {
	fmt.Println("Initialization...")
	connection = "Connect to MySQL"
}

func GetDatabase() string {
	return connection
}