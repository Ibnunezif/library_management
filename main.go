package main

import (
	"BACKEND_GO/controllers"
	"BACKEND_GO/services"
	"BACKEND_GO/models"
)

func main() {
	// Create a new library
	lib := services.NewLibrary()

	// Add some sample members (required to borrow books)
	lib.Members[1] = models.Member{ID: 1, Name: "Alice"}
	lib.Members[2] = models.Member{ID: 2, Name: "Bob"}

	// Start the console interface
	controllers.StartCLI(lib)
}