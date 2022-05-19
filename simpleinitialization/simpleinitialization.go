package main

import "fmt"

func main() {

	if ddLengthMins, cawLengthMins := 250, 30; ddLengthMins > cawLengthMins {
		fmt.Println("docker deep dive is longer than containers on aws wavelength")
		if ddLengthMins > 240 {
			fmt.Println("too much")

		}
	} else if ddLengthMins == cawLengthMins {
		fmt.Println("both have same wavelength")

	} else {
		fmt.Println(" containers on aws wavelength must be longer than docker deep dive")

	}
}
