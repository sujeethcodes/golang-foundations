package main

import (
	"encoding/json"
	"fmt"
)

// Marshal is convert go structure (struct, interface, slice, array, map)to json
// UnMarshal is JSON data convert into Go structure

type UserData struct {
	Name  string
	Age   int
	Email string
}

func main() {
	//--------------
	// Marshal
	user := UserData{Name: "sujeeth", Age: 26, Email: "sujeethcodes@gmail.com"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("Marshal data", string(jsonData))

	//--------------
	// UnMarshal

	json.Unmarshal([]byte(jsonData), &user)
	fmt.Println("UnMarshal data", user)

}
