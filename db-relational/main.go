package main

// import "github.com/go-sql-driver"
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	str := `CREATE TABLE IF NOT EXISTS users (
		id serial,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		age INT,
		primary key (id)
	  );`

	result, err := db.Exec(str)
	if err != nil {
		panic(err)
	}
	_ = result

	// for i := range 10 {
	// 	i++
	// 	_, err = db.Exec(fmt.Sprintf("INSERT INTO users(first_name,last_name,age) VALUES('first_name%d','last_name%d',%d)", i, i, i))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	row := db.QueryRow("SELECT * FROM users WHERE id = 1")
	user := new(User)
	row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age)
	fmt.Println("Your ID:", user.ID)
	fmt.Println("Your first name:", user.FirstName)
	fmt.Println("Your last name:", user.LastName)
	fmt.Println("Your age:", user.Age)

	db.Exec("DELETE FROM users WHERE id = 2")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		user := new(User)
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age)
		fmt.Println("Your ID:", user.ID)
		fmt.Println("Your first name:", user.FirstName)
		fmt.Println("Your last name:", user.LastName)
		fmt.Println("Your age:", user.Age)
		fmt.Println()
	}

}
