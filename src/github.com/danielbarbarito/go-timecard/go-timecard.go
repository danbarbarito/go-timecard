package main

import (
	"fmt"
	"os"
	"io/ioutil"
)


func main() {
	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if (err == nil){
		fmt.Print(string(input))
	}
}
