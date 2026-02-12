package services

import (
	"fmt"
	"sync"
	"time"

	"LIBRARY_MANAGEMENT/models"
	"LIBRARY_MANAGEMENT/concurrency"
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
	Books    map[int]models.Book
	Members  map[int]models.Member
	Mu       sync.Mutex
	ReserveC chan concurrency.ReservationRequest
}

// Factory function
func NewLibrary() LibraryManager {
	Lib := &Library{
		Books:    make(map[int]models.Book),
		Members:  make(map[int]models.Member),
		ReserveC: make(chan concurrency.ReservationRequest, 100),
	}

	go Lib.ReservationWorker() 
	return Lib
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
	fmt.Println("The book is added successfully!")
}

// RemoveBook removes a book by ID
func (l *Library) RemoveBook(bookID int) {
	_, exist := l.Books[bookID]
	if exist {
		delete(l.Books, bookID)
		fmt.Println("The book is deleted successfully!")
	} else {
		fmt.Println("Book not found")
	}
}

// BorrowBook allows a member to borrow a book
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

// ReturnBook allows a member to return a borrowed book
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

// ListAvailableBooks returns all available books
func (l *Library) ListAvailableBooks() []models.Book {
	res := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			res = append(res, book)
		}
	}
	return res
}

// ListBorrowedBooks returns books borrowed by a member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

// ReserveBook enqueues a reservation request
func (l *Library) ReserveBook(bookID int, memberID int) error {
	result := make(chan error)

	req := concurrency.ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Result:   result,
	}

	l.ReserveC <- req
	return <-result
}

// reservationWorker handles reservation requests concurrently
func (l *Library) ReservationWorker() {
	for req := range l.ReserveC {
		go func(r concurrency.ReservationRequest) {
			l.Mu.Lock()
			book, ok := l.Books[r.BookID]
			if !ok {
				r.Result <- fmt.Errorf("Book not found")
				l.Mu.Unlock()
				return
			}

			if book.Status == "Borrowed" || book.Status == "Reserved" {
				r.Result <- fmt.Errorf("Already reserved or borrowed")
				l.Mu.Unlock()
				return
			}

			book.Status = "Reserved"
			l.Books[r.BookID] = book
			l.Mu.Unlock()

			// auto-cancel reservation after 5 seconds if not borrowed
			go func() {
				time.Sleep(5 * time.Second)
				l.Mu.Lock()
				b := l.Books[r.BookID]
				if b.Status == "Reserved" {
					b.Status = "Available"
					l.Books[r.BookID] = b
				}
				l.Mu.Unlock()
			}()

			r.Result <- nil
		}(req)
	}
}