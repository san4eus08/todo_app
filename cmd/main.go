package main

import (
	"github.com/san4eus08/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running server: error code %s", err.Error())
	}
}
