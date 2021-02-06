package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name    string
	address string
}

func main() {
	// RFC - Request for comment
	//URI - Uniform Resource Identifier
	//protocol packages
	//http.Get(url) --> make web clients, server
	// tcp/ip, socket programming
	// net.Dial("tcp", "url: port")
	p1 := Person{name: "Vinh", address: "Boston"}
	fmt.Println(p1)
	barr, err := json.Marshal(p1)
	if err == nil {
		fmt.Println(string(barr))
	}
	var p2 Person
	err = json.Unmarshal(barr, &p2)
}
