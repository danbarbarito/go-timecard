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
	weekday   time.Weekday
	startTime [4]int //index 3 indicates a.m. or p.m.	0=a.m	1=p.m
	duration  float64
	startHour int //Start hour in 12 hours instead of 24 hours
	amOrPm    string
}

//Map to convert strings to time.Weekday
var days = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
}

func generateShifts(filename string) []Shift {
	var shifts []Shift
	//log.Println("Reading file " + filename) TODO - Properly set up logger so it logs to a file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening hours file: ", err)
	}

	scanner := bufio.NewScanner(f) //Scanner used for navigating the file
	shift := Shift{}               //shift struct that will contains the users shift

	for scanner.Scan() {
		var startHour int
		var amOrPm string
		words := strings.Split(scanner.Text(), " ") //Split lines by spaces to access each word
		_, ok := days[strings.Title(words[0])]      //Check if the weekday is correct
		if !ok {
			log.Fatal("Cannot parse hours file: Weekday is incorrect")
		}
		shift.weekday = days[strings.Title(words[0])]        //Store the first word as a time.Weekday
		startTime, err := time.Parse(time.Kitchen, words[1]) //Parse the time
		if err != nil {
			log.Fatal("Cannot parse hours file: ", err)
		}
		shift.startTime[0], shift.startTime[1], shift.startTime[2] = startTime.Clock()
		duration, err := strconv.ParseFloat(words[2], 32) //Parse the duration
		if err != nil {
			log.Fatal("Cannot parse hours file: ", err)
		}
		shift.duration = duration

		//Convert to 12 hours time and set whether it is a.m or p.m
		if shift.startTime[0] > 12 {
			shift.startTime[3] = 1 //Set to p.m
			amOrPm = "p.m"
			startHour = (shift.startTime[0] - 12)
		} else if shift.startTime[0] == 12 {
			shift.startTime[3] = 1 //Set to p.m
			amOrPm = "p.m"
			startHour = (shift.startTime[0])
		} else {
			shift.startTime[3] = 0 //Set to a.m
			amOrPm = "a.m"
			startHour = shift.startTime[0]
		}
		shift.startHour, shift.amOrPm = startHour, amOrPm
		shifts = append(shifts, shift)
	}
	return shifts
}

//Returns the date of when the next weekday begins relative to the 'from' time
func DateOfWeekday(weekday time.Weekday, from time.Time) time.Time {
	var nextDay time.Time
	for i := 0; i < 7; i++ {
		nextDay = from.AddDate(0, 0, i)
		if nextDay.Weekday() == weekday {
			break
		}
	}
	return nextDay
}

//Prints all of the shifts to the stdout
func PrintShifts(shifts []Shift) {
	for i := 0; i < len(shifts); i++ {
		shift := shifts[i]
		fmt.Printf("You work on %s starting at %d:%02d %s for %.1f hours\n", shift.weekday, shift.startHour, shift.startTime[1], shift.amOrPm, shift.duration)
	}
}

//Returns an array of shifts worked from the given time until given duration
func ShiftsWorked(shifts []Shift, from time.Time, until time.Duration) []Shift {
	var shiftsWorked []Shift
	for i := 0; i < len(shifts); i++ {
		shiftsWorked[i] = shifts[i] //The shifts worked are going to be the same as the given shifts except for the
	}
	var _ = days["Sunday"]
	return shiftsWorked
}

func main() {
	//Check if filename is provided
	if len(os.Args) != 2 {
		fmt.Print("Usage: go-timecard <hours file>\n\n\n")
		log.Fatal("Must include a filename")
	}
	filename := os.Args[1]
	shifts := generateShifts(filename) //Store all of the shifts found in the hours file
	PrintShifts(shifts)                //Prints the shifts
	//day := DateOfWeekday(time.Friday, time.Now())
}
