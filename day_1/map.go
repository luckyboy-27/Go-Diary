package main

import "fmt"

func main() {
	//map is golang implementation of hash tables
	// map is a hash table, can use make to create a map
	var id map[string]int     // string is data type of key, int is data type for value (key & value)
	id = make(map[string]int) // create an empty map
	// map litteral
	a := map[string]int{"vinh": 123, "rose": 3} // can have two values to test assignment for existence of key
	//accessing map the same as accessing array, or dict in python
	id["Jack"] = 212  // adding new key-value
	x, y := a["Jack"] // a is false, id is true
	delete(a, "rose")
	fmt.Println(id["Jack"])
	fmt.Println(a)
	fmt.Println(x, y)
	// x is value, y is presence of key, y true if key is in the map
	//iteration
	for key, val := range a {
		fmt.Println(key, val) // print both key and values as pair
	}
}
