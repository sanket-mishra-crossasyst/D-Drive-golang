package main

import (
	"onetomany/databaseconnection"
	"onetomany/handler"
)

func main() {
	databaseconnection.Connection()
	handler.Handler()

}
