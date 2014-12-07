package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)


func main() {
	//Check if filename is provided
	if len(os.Args) == 1{
		fmt.Print("Usage: ./go-timecards <hours file>\n\n\n")
		log.Fatal("Must include a filename")
	}
	fmt.Println("Reading file...")
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text());
	}
}
