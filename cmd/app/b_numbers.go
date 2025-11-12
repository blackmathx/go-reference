package app

import (
	"fmt"
	"math"
)

func Numbers() {
	fmt.Println()
	fmt.Println("------- Numbers ----------------------")
	fmt.Println()

	// Instantiating numbers
	var intNum int = 42                       // Integer
	var floatNum float64 = 3.14               // Floating-point number
	var complexNum complex128 = complex(2, 3) // Complex number

	println(complexNum)

	// Type casting
	_ = float64(intNum) // Casting int to float64
	_ = int(floatNum)   // Casting float64 to int (truncates the decimal part)

	// Common math methods
	fmt.Println("Absolute value:", math.Abs(-42.5))
	fmt.Println("Power (2^3):", math.Pow(2, 3))
	fmt.Println("Square root of 16:", math.Sqrt(16))
	fmt.Println("Ceiling of 3.14:", math.Ceil(3.14))
	fmt.Println("Floor of 3.14:", math.Floor(3.14))
	fmt.Println("Maximum of 10 and 20:", math.Max(10, 20))
	fmt.Println("Minimum of 10 and 20:", math.Min(10, 20))
	fmt.Println("Sine of 90 degrees (in radians):", math.Sin(math.Pi/2))
	fmt.Println("Cosine of 0 degrees (in radians):", math.Cos(0))
	fmt.Println("Tangent of 45 degrees (in radians):", math.Tan(math.Pi/4))
}
