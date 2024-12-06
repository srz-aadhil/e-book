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

	user.Username = "Shambu"
	user.Mail = "shambu99@yrz.com"
	user.Password = "shambu007"
	user.Salt = "randomx123"
	userid, err := user.CreateUser(db)
	if err != nil {
		fmt.Print("User creation failed", err)
	} else {
		fmt.Println("User created with userid- ", userid)
	}

}
