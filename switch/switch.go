package main

import "fmt"

func DayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func main() {
	fmt.Println(DayOfTheWeek(1))
}