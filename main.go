package main

import (
	"ebookmod/pkg/database"
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

	//Creat a user
	user.Username = "babu"
	user.Mail = "babu9@yrz.com"
	user.Password = "babus123"
	user.Salt = "randomx1234"
	userid, err := user.CreateUser(db)
	if err != nil {
		fmt.Println("User creation failed", err)
	} else {
		fmt.Println("User created with userid- ", userid)
	}

	//Get a single user
	oneuser, err := repo.GetUser(db, 7)
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
	// if err = repo.DeleteUser(db, 3); err != nil {
	// 	fmt.Println("User deletion failed due to", err)
	// }

}
