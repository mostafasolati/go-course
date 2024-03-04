package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func main() {

	// for i := range 10 {
	// 	user := &User{
	// 		FirstName: fmt.Sprintf("First Name %d", i+1),
	// 		LastName:  fmt.Sprintf("Last Name %d", i+1),
	// 		Email:     fmt.Sprintf("Email %d", i+1),
	// 	}
	// 	buf := new(bytes.Buffer)
	// 	json.NewEncoder(buf).Encode(user)
	// 	http.Post("http://localhost:5050/user", "application/json", buf)
	// }

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:5050/user/12", nil)
	// req.Header.Set("authorization","dssdfsd")
	http.DefaultClient.Do(req)

	resp, err := http.Get("http://localhost:5050/user")
	if err != nil {
		panic(err)
	}
	users := make([]*User, 0)
	json.NewDecoder(resp.Body).Decode(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}
