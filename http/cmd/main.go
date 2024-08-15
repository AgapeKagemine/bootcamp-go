package main

import (
	"fmt"

	http "github.com/benebobaa/hatetepe"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	router := http.NewRouter()

	router.HandleFunc("POST", "/user", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := r.ParseJSON(&user)
		if err != nil {
			w.WriteHeader(404)
			w.WriteJSON(map[string]string{"message": "ERROR"})
			return
		}
		user.Age += 1
		w.WriteJSON(user)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error")
	}
}
