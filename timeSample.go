package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("This is my current time : ", currentTime.Format("2006-01-02"))
	fmt.Println(currentTime.Format(time.RFC3339Nano)) //predefined time format

	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month())
	fmt.Println("The day is", currentTime.Day())
	fmt.Println("The hour is", currentTime.Hour())
	fmt.Println("The minute is", currentTime.Hour())
	fmt.Println("The second is", currentTime.Second())

	// RFC 3339, the RFC defines a standard format to use for timestamps on the internet.
	//Go time package provides a time.Parse function to convert a string into a time.Time value.
	timeString := "2021-08-15 02:30:45"
	newTime, err := time.Parse("2006-01-02 03:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}

	fmt.Println(newTime)

}
