package main

import (
	"fmt"
	"go-reference/cmd/app"
)

func main() {
	fmt.Println("go-reference")

	app.Intro()
	app.Numbers()
	app.Strings()
	app.Arrays_and_slices()
	app.Maps()

	fmt.Println("end")
}
