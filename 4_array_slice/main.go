package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 5
	arr[1] = 58
	arr[3] = 4
	arr[4] = 2

	// Array'yi ekranda görkezmek üçin fmt.Println
	fmt.Println(arr)

	// double array

	doubleArray := [2][4]int{
		{2, 4, 6, 8},
		{1, 3, 5, 7},
	}

	fmt.Println(doubleArray)
	fmt.Println("array len:", len(doubleArray))

	// Indeksleri 0-dan len(doubleArray)-1-a çenli deňeşdiriň
	for i := 0; i < len(doubleArray); i++ {
		// Ikinji ölçegiň indexini hem dogry dolandyrmak üçin j < len(doubleArray[i])
		for j := 0; j < len(doubleArray[i]); j++ {
			fmt.Print(doubleArray[i][j], " ")
		}
		fmt.Print("\n")
	}

	// slice

	slice := []byte{'a', 'b', 'c', 'd'}

	fmt.Println("slice :", string(slice))

	var array1 = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	var a, b []byte

	a = array1[2:5] // now 'a' has elements ar[2],ar[3] and ar[4]

	b = array1[3:5] // now 'b' has elements ar[3] and ar[4]

	fmt.Println(string(a))
	fmt.Println(string(b))

	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// define two slices
	var aSlice, bSlice []byte

	// some convenient operations
	aSlice = array[:3] // equals to aSlice = array[0:3] aSlice has elements a,b,c
	aSlice = array[5:] // equals to aSlice = array[5:10] aSlice has elements f,g,h,i,j
	aSlice = array[:]  // equals to aSlice = array[0:10] aSlice has all elements

	// slice from slice
	aSlice = array[3:7]  // aSlice has elements d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice contains aSlice[1], aSlice[2], so it has elements e,f
	bSlice = aSlice[:3]  // bSlice contains aSlice[0], aSlice[1], aSlice[2], so it has d,e,f
	bSlice = aSlice[0:5] // slice could be expanded in range of cap, now bSlice contains d,e,f,g,h
	bSlice = aSlice[:]   // bSlice has same elements as aSlice does, which are d,e,f,g

	fmt.Println(string(bSlice))

	// map

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map has two return values. For the second return value, if the key doesn't
	//exist，'ok' returns false. It returns true otherwise.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		rating["C#"] = 3.4
		fmt.Println("We have no rating associated with C# in the map")
		fmt.Println(" but added ")
		fmt.Println(rating)
	}

	delete(rating, "C")
	// deleted C
	fmt.Println("deleted C")
	fmt.Println(rating)
}
