// Package Name
package main

// Import Library
import (
	"errors"
	"fmt"
	"math"
)

// Main Function
func main() {
	// Basic Print Line
	fmt.Println("Hello, World")

	// Basic Variable
	/* var x int = 5
	var y int = 7
	var sum int = x + y
	*/

	// Implicit Variable
	x := 5
	y := 7
	sum := x + y

	fmt.Print("X + Y = ")
	fmt.Println(sum)

	// basic if, else if, else
	if x > y {
		fmt.Println("X more than Y")
	} else if x < y {
		fmt.Println("Y more than X")
	} else {
		fmt.Println("X equals Y")
	}

	// basic array
	// var array [5]int

	// implicit array
	// array := [5]int{0, 0, 0, 0, 0}

	// objecting index array
	array := make(map[string]int)
	array["index-0"] = 0
	array["index-1"] = 1
	fmt.Println(array)

	// basic loops
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// explicit loops
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// array loops
	for index, value := range array {
		fmt.Println("index:", index, "value:", value)
	}

	// basic function
	fmt.Println("X + Y =", sumFunction(x, y))

	// basic error handling function
	// resultSqrt, err := sqrtFunction(10)
	resultSqrt, err := sqrtFunction(-10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSqrt)
	}

	// basic object
	p := person{name: "Andy Maulana Yusuf", age: 20}
	fmt.Println("NAME:", p.name, "AGE", p.age)

	// basic cursor & pointer *
	changeName(&p.name)
	fmt.Println("NAME CHANGE TO:", p.name)
}

func sumFunction(x int, y int) int {
	return x + y
}

func sqrtFunction(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined Negative Number")
	}
	return math.Sqrt(x), nil
}

func changeName(name *string) {
	*name = "AMOEL"
}

type person struct {
	name string
	age  int
}
