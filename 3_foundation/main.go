package main

import (
	"errors"
	"fmt"
	"reflect"
)

// const Pi = 3.14
const Pi float32 = 3.14

func main() {
	var i int = 2
	var isActive bool

	var a int8 = 2
	var b int32 = 12
	// note
	// iki sany int8 we int32 sanlar gosulan yagdayynda error beryar
	// sonun ucin hem in uly type gornusine gecirmeli
	// meselem
	c := int32(a) + b
	fmt.Println(c)
	fmt.Printf("c type : %v", reflect.TypeOf(c))

	isActive = true
	fmt.Println("int number :", i)
	fmt.Println("boolean type :", isActive)

	// strings
	// var s string = "hello"
	// su gornusde alyp bolmayar
	// s[0] = 'c'

	s := "hello"

	// string-daki harpy uytgetjek bolsan byte gornusinde gecirmeli byte = int8
	s1 := []byte(s) // convert string to []byte type

	s1[0] = 'c'
	s2 := string(s1)

	fmt.Println("s2:", s2)
	// iki sany string-i gosmak

	hello := "hello"
	world := " world"

	helloWorld := hello + world
	fmt.Println(helloWorld)

	// s3 := "hello"

	// s3 := "c" + s[1:] // you cannot change string values by index, but you can get values instead.

	// fmt.Printf("%s\n", s)

	m1 := `hello  
		World`
	// yokaradaky gornusde berlende icindaki maglumat hic hili uytgeman yazylyar
	// nahili gornusde yazylan bolsa output hem sol gornusde bolyar
	fmt.Println(m1)

	// error types

	err := errors.New("this is error type")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

}
