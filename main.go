package main

import (
    "LIBRARY_MANAGEMENT/controllers"
    "LIBRARY_MANAGEMENT/services"
    "LIBRARY_MANAGEMENT/models"
)

func main() {
    // Create a new library
    lib := services.NewLibrary()

    // Add sample members directly
    realLib := lib.(*services.Library) 
    realLib.Members[1] = models.Member{ID: 1, Name: "Alice"}
    realLib.Members[2] = models.Member{ID: 2, Name: "Bob"}

    // Add sample books
    realLib.Books[1] = models.Book{ID: 1, Title: "Go Programming", Author: "John"}
    realLib.Books[2] = models.Book{ID: 2, Title: "Algorithms", Author: "Alice"}

    // Start the console interface
    controllers.StartCLI(lib)
}