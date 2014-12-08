package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Shift struct {
	weekday   string
	startTime [4]int // index 3 indicates a.m. or p.m.	0=a.m	1=p.m
	hours     float64
}


//Returns the start hour as a string and a string am or pm
func timeStrings(startTime [4]int)(string, string){
	//Set time to a.m. or p.m.
	var amOrPm = "a.m"
	if startTime[0] > 12 {
		startTime[3] = 1 //Set to p.m.
		amOrPm = "p.m"
	}
	//Convert hour from 24 hours to 12 hours
	var startHour = strconv.FormatInt(int64(startTime[0]), 10)
	if startTime[0] > 12 {
		startHour = strconv.FormatInt(int64(startTime[0]-12), 10)
	}
	return startHour, amOrPm
}

func generateShifts(filename string) []Shift{
	var shifts []Shift
	//log.Println("Reading file " + filename) TODO - Properly set up logger so it logs to a file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	scanner := bufio.NewScanner(f) //Scanner used for navigating the file
	shift := Shift{}               //shift struct that will contains the users shift

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ") //Split lines by spaces to access each word
		shift.weekday = words[0]
		startTime, _ := time.Parse(time.Kitchen, words[1]) //Parse the time
		if err != nil {
			log.Fatal("Cannot parse hours file:", err)
		}
		shift.startTime[0], shift.startTime[1], shift.startTime[2] = startTime.Clock()
		hours, err := strconv.ParseFloat(words[2], 32) //Parse the hours
		if err != nil {
			log.Fatal("Cannot parse hours file:", err)
		}
		shift.hours = hours
		startHour, amOrPm := timeStrings(shift.startTime)
		fmt.Printf("You work on %s starting at %s:%02d %s for %.1f hours\n", shift.weekday, startHour, shift.startTime[1], amOrPm, shift.hours)
		shifts = append(shifts, shift)
	}
	return shifts
}

func main() {
	//Check if filename is provided
	if len(os.Args) == 1 {
		fmt.Print("Usage: ./go-timecards <hours file>\n\n\n")
		log.Fatal("Must include a filename")
	}
	filename := os.Args[1]
	shifts := generateShifts(filename)
	fmt.Print(shifts)
}
