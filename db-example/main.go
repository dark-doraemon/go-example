package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Created_at time.Time `json:"create_at"`
}

func main() {
	connectionString := "host=localhost port=5432 user=postgres password=123456 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect DB successfully")

	rows, err := db.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Age, &u.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Time: %s\n",
			u.Id,
			u.Name,
			u.Created_at.Format("2006-01-02 15:04:05"),
		)
	}

	fmt.Println("JSON format")
	for _, u := range users {
		data, _ := json.Marshal(u)
		fmt.Println(string(data))
	}

}
