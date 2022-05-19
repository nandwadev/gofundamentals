package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

/*var (
	name   = "nandwa"
	course = "Getting started with go"
	module = "10"
	clip   = 2
	test   = 9.2
)*/

func main() {
	name := os.Getenv("USERNAME")
	course := "Getting started with go"
	module := "10"
	clip := 2
	test := 9.2
	courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("module and Clip are set to", module, "and", clip, ".")
	fmt.Println("module is of type", reflect.TypeOf(module))
	fmt.Println("test is of type", reflect.TypeOf(test))
	fmt.Println("course complete is of type", reflect.TypeOf(courseComplete))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("module plus clip equals", total)
	}

	fmt.Println("Memory address of *course* variable is", &course)

	var ptr *string = &course
	fmt.Println("pointing course variable at address,", ptr, "which holds this value,", *ptr)

}
