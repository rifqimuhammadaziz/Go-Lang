package main

import (
	"flag" // golang.org/pkg/flag
	"fmt"
)

func main() {
	// initialize default value
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root",  "Put your database password")
	var port *int = flag.Int("port", 8080, "Put your port application")
	flag.Parse()
	fmt.Println("Host:", *host)
	fmt.Println("Username:", *user)
	fmt.Println("Password:", *password)
	fmt.Println("Port:", *port)
}