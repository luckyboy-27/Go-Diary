package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Name struct {
	fname string
	lname string
}

func main() {
	sli := make([]Name, 0, 100)
	file, err := os.Open("name.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var name Name
		name.fname = s[0]
		name.lname = s[1]
		sli = append(sli, name)
	}
	file.Close()
	fmt.Println(sli)
}
