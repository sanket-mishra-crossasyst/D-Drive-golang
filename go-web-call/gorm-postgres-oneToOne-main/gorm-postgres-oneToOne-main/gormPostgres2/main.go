package main

import (
	"demo/connection"
	"demo/handlers"
)

func main() {
	connection.DatabaseConnection()
	handlers.Controller()
}
