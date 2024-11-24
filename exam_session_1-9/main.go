package main

import (
	"errors"
	"fmt"
)

const maxBorrowLimit = 3

type Book struct {
	BookID     int
	Title      string
	Author     string
	IsBorrowed bool
}

type User struct {
	UserID        int
	Name          string
	BorrowedBooks []int
}

var books = []Book{
	{1, "The Go Programming Guide", "John Doe", false},
	{2, "Concurrency in Practice", "Jane Smith", false},
	{3, "Microservices Design", "Alan Turing", false},
	{4, "Data Structures", "Grace Hopper", false},
	{5, "The Art of Computer Prog", "Donald Knuth", false},
	{6, "Clean Code", "Robert Martin", false},
	{7, "Algorithms Unlocked", "Thomas Cormen", false},
	{8, "Software Architecture", "Martin Fowler", false},
	{9, "System Design", "Scott Meyers", false},
	{10, "Effective Go", "Rob Pike", false},
}

var users = map[int]User{
	1: {1, "Ali", []int{}},
	2: {2, "Babek", []int{}},
	3: {3, "Cabir", []int{}},
	4: {4, "Diana", []int{}},
	5: {5, "Elvira", []int{}},
}

// Displays all books that are currently available
func listAvailableBooks() {
	fmt.Println("Available Books:")
	for _, book := range books {
		if !book.IsBorrowed {
			fmt.Printf("- [%d] %s by %s\n", book.BookID, book.Title, book.Author)
		}
	}
}

// borrowBook allows a user to borrow a book if conditions are met
func borrowBook(userID, bookID int) error {
	user, exists := users[userID]
	if !exists {
		return errors.New("user not found")
	}

	if len(user.BorrowedBooks) >= maxBorrowLimit {
		return errors.New("borrow limit reached")
	}

	for i := range books {
		if books[i].BookID == bookID {
			if books[i].IsBorrowed {
				return fmt.Errorf("book '%s' is already borrowed", books[i].Title)
			}
			books[i].IsBorrowed = true
			users[userID] = User{
				UserID:        user.UserID,
				Name:          user.Name,
				BorrowedBooks: append(user.BorrowedBooks, bookID),
			}
			return nil
		}
	}
	return errors.New("book not found")
}

// returnBook allows a user to return a borrowed book
func returnBook(userID, bookID int) error {
	user, exists := users[userID]
	if !exists {
		return errors.New("user not found")
	}

	for i := range books {
		if books[i].BookID == bookID {
			if !books[i].IsBorrowed {
				return errors.New("book is not currently borrowed")
			}
			for j, bID := range user.BorrowedBooks {
				if bID == bookID {
					books[i].IsBorrowed = false
					user.BorrowedBooks = append(user.BorrowedBooks[:j], user.BorrowedBooks[j+1:]...)
					users[userID] = user
					return nil
				}
			}
			return errors.New("book not found in user's borrowed list")
		}
	}
	return errors.New("book not found")
}

// generateReport displays the current state of the library
func generateReport() {
	fmt.Println("\nLibrary Report:")
	listAvailableBooks()
	fmt.Println("\nBorrowed Books by Users:")
	for _, user := range users {
		fmt.Printf("- %s: ", user.Name)
		if len(user.BorrowedBooks) == 0 {
			fmt.Println("No books borrowed")
		} else {
			for _, bID := range user.BorrowedBooks {
				for _, book := range books {
					if book.BookID == bID {
						fmt.Printf("%s, ", book.Title)
					}
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	fmt.Println("Welcome to the E-Library Book Rental System!")

	// Initial display of available books
	listAvailableBooks()

	// Borrowing books
	fmt.Println("\nUser 3 borrowing Book 1...")
	if err := borrowBook(3, 1); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nUser 3 borrowing Book 2...")
	if err := borrowBook(3, 2); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nUser 4 borrowing Book 5...")
	if err := borrowBook(4, 5); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nUser 4 borrowing Book 10...")
	if err := borrowBook(4, 10); err != nil {
		fmt.Println("Error:", err)
	}
	// Attempting to borrow an already borrowed book
	fmt.Println("\nUser 2 trying to borrow Book 1...")
	if err := borrowBook(2, 1); err != nil {
		fmt.Println("Error:", err)
	}

	// Returning books
	fmt.Println("\nUser 3 returning Book 1...")
	if err := returnBook(3, 1); err != nil {
		fmt.Println("Error:", err)
	}

	// Final report
	generateReport()
}
