package main

import "fmt"

// goto

func myFunc() {
	i := 0
Here:

	fmt.Println(i)
	i++
	if i == 11 {
		return
	}
	goto Here

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndMultiple(a, b int) (int, int) {
	sum := a + b
	multiple := a * b
	return sum, multiple
}

func add(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 4
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	if x == 4 {
		fmt.Println("The integer is equal to 4")
	} else if x < 4 {
		fmt.Println("The integer is less than 4")
	} else {
		fmt.Println("The integer is greater than 4")
	}
	// goto
	myFunc()

	// for loop
	sum := 0
	for index := 0; index <= 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	for index := 10; index > 0; index-- {
		if index == 5 {
			break // or continue
		}
		fmt.Println(index)
	}

	// switch

	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	// Functions

	x = 3
	y := 4
	z := 5
	fmt.Println("max")
	max_xy := max(x, y) // call function max(x, y)
	max_xz := max(x, z) // call function max(x, z)
	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // call function here

	// function Multi-value return

	sum2, multiple := SumAndMultiple(x, y)
	fmt.Printf("sum : %d + %d = %d \n", x, y, sum2)
	fmt.Printf("multiple :  %d * %d = %d \n", x, y, multiple)

	// Pass by value and pointers

	// egerde bir funkisya baha geciremizde pointer bilen gecirsek onda onun asyl oz bahasy hem uytgeyar
	// meselem

	number := 5

	add := add(&number)

	fmt.Println("add +1 function", add)
	fmt.Println("number", number)

}
