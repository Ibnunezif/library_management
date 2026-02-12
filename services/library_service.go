package services

import (
	"fmt"
	"BACKEND_GO/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
	fmt.Println("The book is added successfully!")
}

func (l *Library) RemoveBook(bookID int) {
	_, exist := l.Books[bookID]

	if exist {
		delete(l.Books, bookID)
		fmt.Println("The book is deleted successfully!")
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("The book not found")
	}

	if book.Status == "Borrowed" {
		return fmt.Errorf("The book is already borrowed!")
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("The member is not found!")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	fmt.Println("You have borrowed successfully!")

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("The library doesn't have this book")
	}

	if book.Status == "Available" {
		return fmt.Errorf("The book was not borrowed")
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("Member not found")
	}

	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("You haven't borrowed this book")
	}

	book.Status = "Available"
	l.Books[bookID] = book

	l.Members[memberID] = member

	fmt.Println("Book returned successfully!")
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	res := []models.Book{}

	for _, book := range l.Books {
		if book.Status == "Available" {
			res = append(res, book)
		}
	}

	return res
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		return []models.Book{}
	}

	return member.BorrowedBooks
}