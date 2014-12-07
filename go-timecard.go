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

func createShifts(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	scanner := bufio.NewScanner(f) //Scanner used for navigating the file
	shift := Shift{}         //shift struct that will contains the users shift

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ") //Split lines by spaces to access each word
		shift.weekday = words[0]
		startTime, _ := time.Parse(time.Kitchen, words[1]) //Parse the time
		if err != nil {
			log.Fatal("Cannot parse hours file:", err)
		}
		shift.startTime[0],shift.startTime[1],shift.startTime[2] = startTime.Clock()
		//Set time to a.m. or p.m.
		var amOrPm = "a.m"
		if(shift.startTime[0] > 12){
			shift.startTime[3] = 1 //Set to p.m.
			amOrPm = "p.m"
		}
		hours, err := strconv.ParseFloat(words[2], 32) //Parse the hours
		if err != nil {
			log.Fatal("Cannot parse hours file:", err)
		}
		shift.hours = hours
		//Convert hour from 24 hours to 12 hours
		var startHour = strconv.FormatInt(int64(shift.startTime[0]),10)
		if(shift.startTime[0] > 12){
			startHour = strconv.FormatInt(int64(shift.startTime[0]-12),10)
		}
		fmt.Printf("You work on %s starting at %s:%02d%s for %.1f hours\n",shift.weekday,startHour,shift.startTime[1],amOrPm,shift.hours)
	}
}

func main() {
	//Check if filename is provided
	if len(os.Args) == 1 {
		fmt.Print("Usage: ./go-timecards <hours file>\n\n\n")
		log.Fatal("Must include a filename")
	}
	filename := os.Args[1]
	fmt.Println("Reading file...")
	createShifts(filename)
}
