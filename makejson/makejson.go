package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	var name string
	var addr string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter address: ")
	fmt.Scan(&addr)
	m["name"] = name
	m["adress"] = addr
	json_object, _ := json.Marshal(m)
	fmt.Println(string(json_object))
}
