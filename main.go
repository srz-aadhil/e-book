package main

import (
	"ebookmod/app/database"
	"ebookmod/repo"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.Initdb()
	if err != nil {
		log.Fatal(err)
	}

	var user repo.User //instance of user

	// Create user
	user.ID = 12
	user.Username = "django"
	user.Mail = "aadhilzzz@yrz.com"
	user.Password = "aaaa"
	user.Salt = "random567"

	userid, err := user.CreateUser(db)
	if err != nil {
		fmt.Println("User creation failed", err)
	} else {
		fmt.Println("User created with userid- ", userid)
	}

	//Get a single user
	oneuser, err := repo.GetUser(db, 13)
	if err != nil {
		fmt.Println("User fetching failed due to-", err)
	} else {
		fmt.Println("User details are", oneuser)
	}

	//Get All Users
	_, err = repo.GetAllUsers(db)
	if err != nil {
		fmt.Println("Fetching failed due to", err)
	}

	//Delete a user
	if err = repo.DeleteUser(db, 13); err != nil {
		fmt.Println("User deletion failed due to", err)
	}

	//Update user
	if err = repo.UpdateUser(db, 12, &user); err != nil {
		fmt.Println("User updation failed due to", err)
	}

	//Author Creation
	var newauthor repo.Author
	newauthor.ID = 8
	newauthor.AuthorName = "Django"
	newauthor.CreatedBy = 13

	id, err := newauthor.CreateAuthor(db)
	if err != nil {
		fmt.Println("Author Creation failed due- ", err)
	} else {
		fmt.Printf("Author created with author id- %d", id)
	}

	//Reading Single Author
	author, err := repo.GetAuthor(db, 9)
	if err != nil {
		fmt.Println("Author fetching failed due to- ", err)
	} else {
		fmt.Printf("Author details with author id '%d' are AuthorName:- %s Createdby:- '%d' CreatedAt:- %v  ", author.ID, author.AuthorName, author.CreatedBy, author.CreatedAt)
	}

	//Reading All Authors
	authors, err := repo.GetAllAuthors(db)
	if err != nil {
		fmt.Println("Authors fetching failed ", err)
	} else {
		fmt.Println(authors)
	}

	//Update an Author
	if err = repo.UpdateAuthor(db,&newauthor); err!= nil {
		fmt.Println("Author updation failed -", err)
	}

	//Delete an Author
	if err = repo.DeleteAuthor(db, 8); err != nil {
		fmt.Println("Author deletion failed due to- ", err)
	}

	//Create a book
	var newbook repo.Book // Instance of book

	newbook.Title = "fresh book"
	newbook.Content = "rare piece version"
	newbook.AuthorID = 5
	newbook.CreatedBy = 9
	newbook.Status = 2 // '1'- Published, '2'- Draft, '3'- Deleted

	 bookid, err := newbook.CreateBook(db)
	if err != nil {
		fmt.Println("Book creation failed due to ", err)
	} else {
		fmt.Printf("Book created with id '%v'", bookid)
	}

	//Get a Book
	book, err := repo.GetaBook(db, 15)
	if err != nil {
		fmt.Printf("Book fetching failed due to- %s", err)
	} else {
		fmt.Printf("Book with id '%d' details are\nTitle:- %s\nContent:- %s\nAuthor:- %d\nCreatedby:- %d\nCreated at:- %v\n", book.ID, book.Title, book.Content, book.AuthorID, book.CreatedBy, book.CreatedAt)
	}

	//Get All books
	allbooks, err := repo.GetAllBooks(db)
	if err != nil {
		fmt.Println("All Books fetching failed due to ", err)
	} else {
		for _, book := range allbooks {
			fmt.Printf("Book with id '%d' details are\nTitle:- %s\nContent:- %s\nAuthor:- %d\nCreatedby:- %d\nCreated at:- %v\n\n", book.ID, book.Title, book.Content, book.AuthorID, book.CreatedBy, book.CreatedAt)
		}
	}

	//Update a book
	if err := repo.UpdateABook(db, &newbook); err != nil {
		fmt.Println("Book updation failed\n", err)
	}

	//Delete a book
	if err := repo.DeleteaABook(db, 6); err != nil {
		fmt.Println("Book deletion failed due to ", err)
	}

}
