package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "ebookdb"
)

var db *sql.DB
var err error

// User CRUD
// User creation function
// func CreateUser(mail, username, password, salt string) (userId int, err error) {
// 	query := `INSERT INTO users  (mail, username, password, salt) VALUES($1,$2,$3,$4) RETURNING id`

// 	if err = db.QueryRow(query, mail, username, password, salt).Scan(&userId); err != nil {
// 		fmt.Printf("User creation failed due to %v", err)
// 	}
// 	return userId, err
// }

// Getting one user details function
// func GetUser(id int) (mail, username string, createdat, updatedat time.Time, err error) {
// 	query := `SELECT mail,username,created_at,updated_at FROM users WHERE id=$1`

// 	if err = db.QueryRow(query, id).Scan(&mail, &username, &createdat, &updatedat); err != nil {
// 		log.Printf("User details fetching failed due to %s", err)
// 	}

// 	fmt.Println("User fetching successfull")
// 	return mail, username, createdat, updatedat, nil
// }

// Getting all users fucntion
// func GetAllUsers() ()

// Update user function
func UpdateUser(id int, mail, password string) (err error) {
	query := `UPDATE users SET mail=$1,password=$2,updated_at=$3 WHERE id =$4`

	result, err := db.Exec(query, mail, password, time.Now().UTC(), id)
	if err != nil {
		return fmt.Errorf("User Updation failed due to %v", err)
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to : %w", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no user with ID : %d", id)
	}

	log.Println("User updation successfully completed")
	return nil
}

func main() {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode= disable", user, password, host, port, dbname)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection successfully established")

	//User creation
	// userId, err := CreateUser("user2@srz.com", "adam11", "user008", "random789")
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// fmt.Println("User created with userId-", userId)

	//Getting one user
	// mail, username, createdat, updatedat, err := GetUser(1)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// fmt.Printf("User details are \n mail-%s, username-%s, createdat-%v, updatedat-%v", mail, username, createdat, updatedat)

	//Updating a user
	if err := UpdateUser(1, "userupdated@srz.com", "newpassword"); err != nil {
		log.Println(err)
	}
}
