package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//open to open a file
	//read: read the bytes into []byte
	//write: write []byte into file
	//close: close file
	//seek: move read/write file
	//dat, e := ioutil.ReadFile("test.txt")
	//large file causes error
	dat := []byte("Hello,World!\n")
	err := ioutil.WriteFile("hello.txt", dat, 0777)
	check(err)

	//way 2
	s := "Hello world"
	f, err := os.Create("hi.txt")
	check(err)
	barr := []byte{1, 2, 3}
	nb, err := f.Write(barr)
	a, err := f.WriteString(s)
	fmt.Println(nb)
	fmt.Println(a)
}
