package main

import "fmt"

func main() {

	name := "nandwa"
	course := "Getting started with go"
	fmt.Println("\nHi", name, "your current course is", course)

	updateCourse(&course)

	fmt.Println("You are currently watching", course)

}

func updateCourse(course *string) string {
	*course = "Getting started with docker and kubernetes"
	fmt.Println("updated course to", *course)

	return *course

}
