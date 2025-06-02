package main

import "fmt"

func main2() {
	day := "sunday"

	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Printf("%s is a workday\n", day)
	default:
		fmt.Printf("%s is a weekend day\n", day)
	}
}
