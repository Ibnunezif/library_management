package controllers

import (
	"fmt"
	"BACKEND_GO/models"
	"BACKEND_GO/services"
)

func StartCLI(l services.LibraryManager) {
	for {
		fmt.Println("\n--- Library Management ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBookCLI(l)
		case 2:
			removeBookCLI(l)
		case 3:
			borrowBookCLI(l)
		case 4:
			returnBookCLI(l)
		case 5:
			listAvailableBooksCLI(l)
		case 6:
			listBorrowedBooksCLI(l)
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addBookCLI(l services.LibraryManager) {
	var id int
	var title, author string

	fmt.Print("Enter book ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter book Title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book Author: ")
	fmt.Scanln(&author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	l.AddBook(book)
}

func removeBookCLI(l services.LibraryManager) {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scanln(&id)
	l.RemoveBook(id)
}

func borrowBookCLI(l services.LibraryManager) {
	var bookID, memberID int
	fmt.Print("Enter your Member ID: ")
	fmt.Scanln(&memberID)
	fmt.Print("Enter Book ID to borrow: ")
	fmt.Scanln(&bookID)

	err := l.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func returnBookCLI(l services.LibraryManager) {
	var bookID, memberID int
	fmt.Print("Enter your Member ID: ")
	fmt.Scanln(&memberID)
	fmt.Print("Enter Book ID to return: ")
	fmt.Scanln(&bookID)

	err := l.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func listAvailableBooksCLI(l services.LibraryManager) {
	books := l.ListAvailableBooks()
	fmt.Println("\nAvailable Books:")
	for _, b := range books {
		fmt.Printf("ID:%d | Title:%s | Author:%s\n", b.ID, b.Title, b.Author)
	}
}

func listBorrowedBooksCLI(l services.LibraryManager) {
	var memberID int
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	books := l.ListBorrowedBooks(memberID)
	fmt.Println("\nBorrowed Books:")
	for _, b := range books {
		fmt.Printf("ID:%d | Title:%s | Author:%s\n", b.ID, b.Title, b.Author)
	}
}