package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// specifiedTime := time.Date(1990, 10, 01, 18, 13, 16, 99, time.Local)

	// Uncomment the following line see specifiedTime in raw format.
	// It will include a lot of information about the time, including nanoseconds and local timezone.
	// fmt.Println(specifiedTime)

	// Uncomment the following line see specifiedTime in the format of year-month-day-hour-minute-second.
	// fmt.Println(specifiedTime.Format("2006-01-02 15:04:05"))

	// Uncomment the following line see specifiedTime in 12 hour format.
	// fmt.Println(specifiedTime.Format("2006-01-02 3:04:05"))

	// Uncomment the following line see specifiedTime without the first two digits of the year, and without the time of day.
	// fmt.Println(specifiedTime.Format("06-1-02"))

	// Uncomment the following line see specifiedTime in the format of day-month-year-hour-minute-second.
	// fmt.Println(specifiedTime.Format("02-01-2006 5:04:05"))

	// Uncomment the following line see specifiedTime in the RFC3339 format.
	// fmt.Println(specifiedTime.Format(time.RFC3339))

	// Uncomment the following line see specifiedTime in the RFC3339Nano format.
	// fmt.Println(specifiedTime.Format(time.RFC3339Nano))

	// WORKING WITH time.Parse

	// Next I will be working with time.Parse to parse a string into a time.Time.
	// From my understanding so far parsing means to turn a string into another data type.
	// In the below case, the other data type is time.Time. time.Time is a struct, (a custom data type).
	// Notice that the time.Parse method accepts two arguments.
	// First argument is the layout to be used, second is the string to be parsed.
	// I'll choose the same date and time as specifiedTime, but I'll use a different format.

	// The below string variable is written in a the RFC3339Nano format.
	// stringTime := "1916-04-24T9:00:00.0000032-05:00"

	// Uncomment the following 5 lines to see how time.Parse works with custom layouts (as opposed to pre-defined layouts).
	// theTime, err := time.Parse("2006-01-02 15:04:05", stringTime)
	// if err != nil {
	// 	fmt.Println("Could not parse the time.", err)
	// }
	// fmt.Println("The time is:", theTime)

	// Uncomment the following 6 lines to see how time.Parse works with the RFC3339Nano layout.
	// theTime, err := time.Parse(time.RFC3339Nano, stringTime)
	// if err != nil {
	// 	fmt.Println("Could not parse the time.", err)
	// }
	// fmt.Println("The time is:", theTime)
	// fmt.Println(theTime.Format(time.RFC3339Nano))

	// Uncomment the following 6 lines to see how I'm using the time.RFC3339 constant as a layout in the time.Parse method.
	// And that I'm using time.Kitchen as the format. time.Kitchen is a predefined layout that only shows the hour and minute in AM or PM.

	// theTime, err := time.Parse(time.RFC3339, stringTime)
	// if err != nil {
	// 	fmt.Println("Could not parse the time.", err)
	// }
	// fmt.Println("The time is:", theTime)
	// fmt.Println(theTime.Format(time.Kitchen))

	// WORKING WITH TIME ZONES.

	// Next I will be working with time zones.
	// The .UTC method returns t with the location set to UTC.
	// In the above line, "t" refers to the variable the .UTC method is called on.
	// So in the below case, it is called on the variable "theTime".
	// It would appear that .UTC does not take any arguments.

	// theTime := time.Date(1916, 4, 24, 9, 30, 0, 0, time.Local)
	// fmt.Println("The time is:", theTime)
	// fmt.Println(theTime.Format(time.RFC3339Nano))

	// Below I am converting the time to UTC using the .UTC() method.

	// utcTime := theTime.UTC()
	// fmt.Println("The UTC time is: ", utcTime)
	// fmt.Println(utcTime.Format(time.RFC3339Nano))

	// Below I am converting the time back to local time using the .Local() method.

	// localTime := utcTime.Local()
	// fmt.Println("The local time is: ", localTime)
	// fmt.Println(localTime.Format(time.RFC3339Nano))

	// WORKING WITH COMPARING TIMES

	// To compare times, I will be using the .Before() , .After() and .Equal() methods.
	// These methods return a boolean true or false whether the comparison is true or false.

	// Supposedly, the .Equal() method can be used to compare two times in different time zones.
	// so 14:00:00.Equal(16:00:00 +2) would return true.

	// firstTime := time.Date(1916, 4, 24, 9, 30, 0, 0, time.UTC)
	// fmt.Println("The first time is: ", firstTime)

	// secondTime := time.Date(1916, 4, 29, 17, 30, 0, 0, time.UTC)
	// fmt.Println("The second time is: ", secondTime)

	// thirdTime := time.Date(1916, 4, 24, 9, 30, 0, 0, time.UTC)
	// fmt.Println("The third time is: ", thirdTime)

	// fmt.Println("First time before second:", firstTime.Before(secondTime))
	// fmt.Println("First time after second:", firstTime.After(secondTime))

	// fmt.Println("Second time before first:", secondTime.Before(firstTime))
	// fmt.Println("Second time after first:", secondTime.After(firstTime))

	// fmt.Println("Is third equal to first?", firstTime.Equal(thirdTime))
	// fmt.Println("Is third equal to second?", secondTime.Equal(thirdTime))

	// WORKING WITH DURATIONS

	// The .Sub() method returns a duration between two times int the form of a time.Duration type.
	// The time.Duration type is a custom data type, and its maximum tim value is one hour.
	// The reason for this max time value of one hours is because some months have varying amounts of days. Daylight savings also affects the time.
	// So it makes sense to have a max time value of one hour.

	// firstTime := time.Date(1916, 4, 24, 9, 30, 0, 0, time.UTC)
	// fmt.Println("The first time is: ", firstTime)

	// secondTime := time.Date(1916, 4, 29, 17, 30, 0, 0, time.UTC)
	// fmt.Println("The second time is: ", secondTime)

	// fmt.Println("Duration between first and second:", firstTime.Sub(secondTime))
	// fmt.Println("Duration between second and first:", secondTime.Sub(firstTime))

	// WORKING WITH ADDING AND SUBTRACTING TIMES

	// The .Add() and .Sub() methods can be used to add or subtract a duration from a time. (Thanks GitHub Copilot)

	// toAdd := 1 * time.Hour
	// fmt.Println("I just created time.Duration variable of 1 hour. This is it:", toAdd)

	// toAdd += 2 * time.Second
	// fmt.Println("I just added 2 seconds to the time.Duration variable. This is it:", toAdd)

	// toAdd += 3 * time.Minute
	// fmt.Println("I just added 3 minutes to the time.Duration variable. This is it:", toAdd)

	// toAdd -= 3*time.Minute + 2*time.Second
	// fmt.Println("I just deducted 3 minutes and 2 seconds from the time.Duration variable. It should be back to 1 hour exactly:", toAdd)

	// Below I am using the .Add() method to add one day (24 * 1 hour) and also one year (365 * 24 * 1 hour) to "specifiedTime".

	// specifiedTime := time.Date(1916, 4, 24, 9, 30, 0, 0, time.UTC)
	// addOneDay := 24 * time.Hour
	// addOneYear := 365 * 24 * time.Hour

	// fmt.Println(specifiedTime)
	// fmt.Println(specifiedTime.Add(addOneDay))
	// fmt.Println(specifiedTime.Add(addOneYear))

	// NOW THE FINAL CHAPTER.
	// COUNTING DOWN THE DAYS LEFT BEFORE A BIRTHDAY.

	// I will be using the .Sub() method to count down the days left before a birthday.
	// t represents the birthday, and u represents the current time.

	// Below is the hardcoded birthday.

	// sarahsBirthday := time.Date(2022, 6, 10, 0, 0, 0, 0, time.UTC)
	// daysRemaining2 := sarahsBirthday.Sub(time.Now())

	// fmt.Println("Sarah's birthday is:", sarahsBirthday.Format(time.RFC850))
	// fmt.Println("Hours until Sarah's birthday remaining:", daysRemaining2)

	// Below is the user inputted birthday.

	var birthday string

	fmt.Println("Please enter your birthday in the format DD-MM-YYYY.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	birthday = scanner.Text()

	parsedDate, err := time.Parse("02-01-2006", birthday)
	if err != nil {
		fmt.Println("There was an error parsing your birthday. Please try again.")
	}

	fmt.Println("The birthday is", parsedDate.Format("02-01-2006"))

	daysRemaining := parsedDate.Sub(time.Now())

	fmt.Println("Hours until your birthday:", daysRemaining)

}
