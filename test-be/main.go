package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type customer struct {
	name    string
	id      uint64
	address string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is My Website!\n")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	c := customer{
		name:    "Raymond",
		id:      2301856485,
		address: "Tangerang",
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["Nama"] = c.name
	resp["Nim"] = fmt.Sprintf("%d", c.id)
	resp["Alamat"] = c.address
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error on JSON Marshal with Err: " + err.Error() + "\n")
	}

	w.Write(jsonResp)
}

func main() {
	var PORT int = 5000
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/semuadata", getAllCustomer)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error Starting Server: %s\n", err)
		os.Exit(1)
	}
}
