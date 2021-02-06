package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Printf("Please input name: ")
	fmt.Scan(&name)
	fmt.Printf("Please input address: ")
	fmt.Scan(&address)
	data := map[string]string{"name": name, "address": address}
	barr, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(barr))
	}
}
