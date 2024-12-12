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

	//var user repo.User //instance of user

	//Creat a user
	// user.ID = 12
	// user.Username = "django"
	// user.Mail = "aadhilzzz@yrz.com"
	// user.Password = "aaaa"
	// user.Salt = "random567"

	// userid, err := user.CreateUser(db)
	// if err != nil {
	// 	fmt.Println("User creation failed", err)
	// } else {
	// 	fmt.Println("User created with userid- ", userid)
	// }

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
	var author repo.Author
	author.AuthorName = "Akhilan"
	author.CreatedBy = 12

	id, err := author.CreateAuthor(db)
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

}
