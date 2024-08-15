package main

import (
	"encoding/json"
	"fmt"

	"marshall/internal/domain"
)

func unmarshall(data []byte) (result domain.User) {
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error: ", err)
		return domain.User{}
	}
	return result
}

func marshall(user domain.User) []byte {
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	return jsonData
}

func main() {
	user := domain.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   30,
	}

	serialized := marshall(user)
	fmt.Println(serialized)

	for _, v := range serialized {
		fmt.Printf("%c", rune(v))
	}
	fmt.Println()

	unserialized := unmarshall(serialized)
	fmt.Println(unserialized)
}
