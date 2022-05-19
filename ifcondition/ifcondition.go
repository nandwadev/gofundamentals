package main

import "fmt"

func main() {
	ddLengthMins := 300
	cawLengthMins := 300

	if ddLengthMins > cawLengthMins {
		fmt.Println("docker deep dive is longer than containers on aws wavelength")
	} else if ddLengthMins == cawLengthMins {
		fmt.Println("both have same wavelength")

	} else {
		fmt.Println(" containers on aws wavelength must be longer than docker deep dive")

	}
}
