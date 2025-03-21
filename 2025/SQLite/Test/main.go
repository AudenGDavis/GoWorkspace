package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mattn/go-sqlite3"
)

func main() {
	// Open a database connection
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("DROP TABLE users")

	db.Exec("CREATE TABLE users (id INT PRIMARY KEY, name INTEGER, age INT)")

	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES (2480, 'Hayden Smith', 25), (2434, 'Taylor Baker', 18), (1429, 'Charlie Jackson', 28), (2491, 'Dallas Torres', 34), (3162, 'Jordan Hall', 20), (6174, 'Cameron Mitchell', 20), (8842, 'Alex Brown', 24), (6113, 'Dakota Rodriguez', 31), (5352, 'Hunter Jones', 32), (3008, 'Kennedy Clark', 39), (1463, 'Shawn Ramirez', 26), (2724, 'Casey Perez', 21), (3223, 'Morgan Roberts', 36), (9231, 'Drew Jackson', 34), (3272, 'Leslie Williams', 30), (3949, 'Reese Adams', 38), (7284, 'Logan White', 37), (8230, 'Corey Baker', 28), (1487, 'Alex Hill', 23), (7072, 'Micah Taylor', 28), (2703, 'Marley Anderson', 18), (7756, 'Emerson Scott', 17), (3736, 'Kennedy Roberts', 37), (3233, 'Chris Hall', 25), (3631, 'Marley Adams', 21), (1919, 'Alex Hill', 24), (9322, 'Kendall Hernandez', 34), (4081, 'Blake Gonzalez', 35), (5874, 'Hunter Wright', 19), (4620, 'Cameron Walker', 16), (6608, 'Jamie Harris', 24), (9662, 'Rowan Johnson', 32), (8759, 'Finley Scott', 25), (2069, 'Kelly Brown', 27), (1031, 'Rowan Carter', 22), (6177, 'Shawn Baker', 35), (2500, 'Sydney Hill', 22), (8807, 'Shawn Hernandez', 30), (5981, 'Parker Sanchez', 36), (1490, 'Reagan Walker', 28), (6729, 'Marley Clark', 27), (3047, 'Corey Carter', 22), (9558, 'Emerson Martin', 31), (8410, 'Kelly Hall', 37), (3720, 'Jamie Scott', 29), (3650, 'Lane Thompson', 28), (4409, 'Micah Hernandez', 18), (2087, 'Leslie Lee', 21), (6080, 'Quincy Anderson', 19), (6783, 'Pat Thomas', 17)")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, age FROM users WHERE age >= 30")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Println("User:", id, name, age)
	}

}
