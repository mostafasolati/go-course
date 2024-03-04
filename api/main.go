package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func createUserHandler(service UserInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(User)
		json.NewDecoder(r.Body).Decode(user)
		fmt.Println("USER:", user)
		err := service.Create(user)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "ERROR!")
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func updateUserHandler(service UserInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(User)
		json.NewDecoder(r.Body).Decode(user)
		fmt.Println("USER:", user)
		err := service.Update(user)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "ERROR!")
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func findUserHandler(service UserInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		s := r.PathValue("id")
		id, _ := strconv.Atoi(s)
		user, _ := service.Find(id)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func listUsersHandler(service UserInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := service.List()
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func deleteUserHandler(service UserInterface) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		s := r.PathValue("id")
		id, _ := strconv.Atoi(s)
		service.Delete(id)
		// w.Header().Set("content-type", "application/json")
	}

}

func main() {
	service := NewUserService()
	// service = ...
	fmt.Println("Server Started...")
	http.HandleFunc("GET /user", listUsersHandler(service))
	http.HandleFunc("POST /user", createUserHandler(service))
	http.HandleFunc("PUT /user", updateUserHandler(service))
	http.HandleFunc("GET /user/{id}", findUserHandler(service))
	http.HandleFunc("DELETE /user/{id}", deleteUserHandler(service))
	http.ListenAndServe("localhost:5050", nil)

	// service := NewUserService()
	// user := &User{
	// 	FirstName: "Sima",
	// 	LastName:  "Nemati",
	// 	Email:     "sima@gmail.com",
	// }

	// service.Create(user)

	// user = &User{
	// 	FirstName: "Arta",
	// 	LastName:  "Alijan",
	// 	Email:     "arta.alijan@gmail.com",
	// }

	// service.Create(user)

	// fmt.Println(service.Find(1))
	// fmt.Println(service.Find(2))
	// service.Delete(2)
	// fmt.Println(service.List())

}
