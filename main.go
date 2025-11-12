package main

import (
	"fmt"
	"go-reference/pkg"
)

func main() {
	fmt.Println("go-reference")

	pkg.Intro()
	pkg.Numbers()
	pkg.Strings()
	pkg.Arrays_and_slices()
	pkg.Maps()

	fmt.Println("end")
}
