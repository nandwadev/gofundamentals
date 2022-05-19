package main

import (
	"fmt"
	"strings"
)

func main() {

	author := "nigel poulton"
	course := "Go Fundamentals"
	fmt.Println(TestFunction(author, course))
}
func TestFunction(author, course string) (authorName, courseName string) {
	author = strings.ToUpper(author)
	course = strings.ToUpper(course)

	return author, course
}
