package db

import (
	"database/sql"
	"fmt"
	"log"
)


func ConnectDB(dsn string) *sql.DB{
	db , err := sql.Open("mysql" , dsn)

	if err != nil {
		log.Fatalf("Error while connecting DB : %v" , err)
	}

	if err := db.Ping(); err != nil{
		log.Fatalf("Error Ping : %v" , err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users(
	id INT AUTO_INCREMENT PRIMARY KEY,
	name varchar(30) not null,
	email varchar(50) not null unique,
	password varchar(100),
	dob DATE
	)
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table users %v :" , err)
	}
	fmt.Println("UserTable created successfully")

	query4 := `
	CREATE TABLE IF NOT EXISTS posts(
	tweet VARCHAR(100),
	userid INT NOT NULL,
	useremail VARCHAR(50) NOT NULL
	)
	`
	_, err = db.Exec(query4)
	if err != nil {
		log.Fatalf("Error creating table POSTS %v :" , err)
	}
	fmt.Println("PostTable created successfully")


	query1 := `
	CREATE TABLE IF NOT EXISTS likes(
	likecount INT AUTO_INCREMENT PRIMARY KEY,
	likebyemail varchar(50) not null
	)
	`
	_, err = db.Exec(query1)
	if err != nil {
		log.Fatalf("Error creating table likes %v :" , err)
	}
	fmt.Println("LikesTable created successfully")

	query2 := `
	CREATE TABLE IF NOT EXISTS comments(
	opinion VARCHAR(100),
	commentersemail VARCHAR(50) not null
	)
	`
	_, err = db.Exec(query2)
	if err != nil {
		log.Fatalf("Error creating table comments %v :" , err)
	}
	fmt.Println("CommentsTable created successfully")

	fmt.Println("Connected to db")
	return db
}