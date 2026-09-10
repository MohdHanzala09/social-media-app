package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error while connecting DB : %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error Ping : %v", err)
	}

	query := `
	CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table users %v :", err)
	}
	fmt.Println("UserTable created successfully")

	query4 := `
	CREATE TABLE tweets (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	content VARCHAR(280) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
	`
	_, err = db.Exec(query4)
	if err != nil {
		log.Fatalf("Error creating table POSTS %v :", err)
	}
	fmt.Println("PostTable created successfully")

	query1 := `
	CREATE TABLE likes (
	id INT AUTO_INCREMENT PRIMARY KEY,
	tweet_id INT NOT NULL,
	user_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UNIQUE KEY unique_like (tweet_id, user_id),
	FOREIGN KEY (tweet_id) REFERENCES tweets(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
	`
	_, err = db.Exec(query1)
	if err != nil {
		log.Fatalf("Error creating table likes %v :", err)
	}
	fmt.Println("LikesTable created successfully")

	query2 := `
	CREATE TABLE comments (
	id INT AUTO_INCREMENT PRIMARY KEY,
	tweet_id INT NOT NULL,
	user_id INT NOT NULL,
	content VARCHAR(500) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (tweet_id) REFERENCES tweets(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
	`
	_, err = db.Exec(query2)
	if err != nil {
		log.Fatalf("Error creating table comments %v :", err)
	}
	fmt.Println("CommentsTable created successfully")

	fmt.Println("Connected to db")
	return db
}
