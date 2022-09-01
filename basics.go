package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
	// level 1 - variables, types, values
	// short declaration operator :=  and single line print
	exercise1()

	// zero values for variables
	exercise2()

	// fmt.Sprintf to print all of the VALUES to one single string.
	exercise3()

	// Create your own type
	exercise4()

	// level 2 - programming fundamentals
	// decimal, binary, and hex + shift operators
	exercise5()

	// TYPED and UNTYPED constants + iota
	exercise6()

	// raw string literal
	exercise7()

	// level 3
	// control flow - for
	exercise8()

	// if elseif else - FizzBuzz + switch case
	exercise9()

	// Array + slice
	exercise10()

	// Map
	exercise11()

	// structs + anonymous structs
	exercise12()

	// structs TODO
	exercise13()

}

func exercise13()  {
	// ● Create a new type: vehicle.
	//○ The underlying type is a struct.
	//○ The fields:
	//■ doors
	//■ color
	//● Create two new types: truck & sedan.
	//○ The underlying type of each of these new types is a struct.
	//○ Embed the “vehicle” type in both truck & sedan.
	//○ Give truck the field “fourWheel” which will be set to bool.
	//○ Give sedan the field “luxury” which will be set to bool. solution
	//● Using the vehicle, truck, and sedan structs:
	//○ using a composite literal, create a value of type truck and assign values to the
	//fields;
	//○ using a composite literal, create a value of type sedan and assign values to the
	//fields.
	//● Print out each of these values.
	//● Print out a single field from each of these values


}
func exercise12(){
	// Create your own type “person” which will have an underlying type of “struct” so that it can store
	//the following data:
	// first name
	// last name
	// favorite ice cream flavors
	//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
	//which stores the favorite flavors.
	type person struct{
		first_name string
		last_name string
		fav_book []string
	}

	p1 := person{
		first_name: "Deepti",
		last_name: "Shukla",
		fav_book: []string{"ABCDEs"},
	}

	fmt.Print(p1.first_name)
	fmt.Print(p1.last_name)

	// keeping structs in map
	m := map[string]person{
		"first_person" : p1,
	}

	fmt.Print(m)

	// anonymous structs
	p := struct{
		first_name string
		last_name string
		fav_book []string
	}{
		first_name: "Deepti",
		last_name: "Shukla",
		fav_book: []string{"ABCDEs"},
	}

	fmt.Print(p.first_name)
}

func exercise11(){
	// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
	//TYPE []string which stores their favorite things. Store three records in your map. Print out all of
	//the values, along with their index position in the slice.
	//`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	//`money_miss`, `James Bond`, `Literature`, `Computer Science`
	//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

	m := map[string][]string{
		`bond_james`: [] string {`Shaken, not stirred`, `Martinis`, `Women`},
		`money_miss`: [] string {`James Bond`, `Literature`, `Computer Science`},
		`no_dr` : [] string { `Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(m)

	// add to map
	m[`fleming_bird`] = [] string {`flying bird`, `flying duck`}

	for k, v := range m {
		fmt.Print("for key:",k,"=>")
		for i, v2 :=  range v{
			fmt.Print(i,"-",v2,",")
		}
		fmt.Println()
	}
	fmt.Println()

	// delet from map
	delete(m, `fleming_bird`)
	for k, v := range m {
		fmt.Print("for key:",k,"=>")
		for i, v2 :=  range v{
			fmt.Print(i,"-",v2,",")
		}
		fmt.Println()
	}
}

func exercise10(){
	// create an ARRAY which holds 5 VALUES of TYPE int
	// assign VALUES to each index position.
	// Range over the array and print the values out.
	// Using format printing
	// print out the TYPE of the array

	x := [5]int{1,2,3,4,5}
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("Type of x:", reflect.TypeOf(x))

	// use slice instead of array
	y := []int{1,2,3,4,5}
	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println("Type of y:", reflect.TypeOf(y))

	// Using the code from the previous example, use SLICING to create the following new slices
	//which are then printed:
	// [42 43 44 45 46]
	// [47 48 49 50 51]
	// [44 45 46 47 48]
	// [43 44 45 46 47]
	z := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(z[:5])
	fmt.Println(z[5:])
	fmt.Println(z[2:7])
	fmt.Println(z[1:6])

	// append slices
	slice_x := []int{1,2,3,4}
	slice_x = append(slice_x, 5,6)
	fmt.Println(slice_x)

	// delete from slice get multiple slices
	//  start with this slice
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// use APPEND & SLICING to get these values here which you should ASSIGN to a
	//variable “y” and then print:
	// [42, 43, 44, 48, 49, 50, 51]
	xyz := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xyz = append( xyz[:3], xyz[6:]...)
	fmt.Println(xyz)


	// 2D array ( slice of slice )
	str1 := []string{}
	str2 := []string{}

	str3 := [][]string{str1, str2}
	fmt.Println(str3)

	seperator()
}

func exercise9(){
	// print Fizz if number is divisible by 3,
	// print Buzz if number is divisible by 5,
	// print FizzBuzz if divisible by both

	for i:=0; i< 10; i++{
		if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println("FizzBuzz")
		}
	}

	// using some switch cases
	a := true
	switch a {
	case false :
		println("should not print")

	case true:
		println("should print")
	}
	seperator()
}

func exercise8(){
	//Print out the remainder (modulus) which is found for each number between 0-15 when it
	//is divided by 4.

	// standard for loop
	for i:=0; i<15; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}

	// for condition {}
	i := 5
	for i > 0 {
		i--;
	}

	// for { condition }
	i = 0
	for {
		if i > 5 {
			break;
		}
		i++;
	}

	seperator()
}

func exercise7(){
	// Create a variable of type string using a raw string literal. Print it

	a := ` This is a
	test for
	raw
	String "literals"`

	fmt.Println(a)

	seperator()
}

func exercise6(){
	// Create TYPED and UNTYPED constants. Print the values of the constants.
	const(
		a = 5
		b int = 5
		)
	fmt.Println(a)
	fmt.Println(b)

	// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	const(
		year1 = 1995 + iota
		year2 = 1995 + iota
		year3 = 1995 + iota
	)
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)

	seperator()
}

func exercise5(){
	// Write a program that
	// assigns an int to a variable
	// prints that int in decimal, binary, and hex
	// shifts the bits of that int over 1 position to the left, and assigns that to a variable
	// prints that variable in decimal, binary, and hex

	var num int64

	num = 10

	fmt.Println("%.2f", num)

	binary_num := strconv.FormatInt(num, 2)
	fmt.Println("hex value:",binary_num)

	hex_num := strconv.FormatInt(num, 16)
	fmt.Println("hex value:",hex_num)

	// shift
	num = num << 1

	fmt.Println("%.2f", num)

	binary_num = strconv.FormatInt(num, 2)
	fmt.Println("hex value:",binary_num)

	hex_num = strconv.FormatInt(num, 16)
	fmt.Println("hex value:",hex_num)
	seperator()
}

func exercise4(){
	//1. Create your own type "year". Have the underlying type be an int.
	//2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
	//a. print out the value of the variable “x”
	//b. print out the type of the variable “x”
	//c. assign 1999 to the VARIABLE “x” using the “=” OPERATOR

	type year int

	var x year
	x = 1999

	//i. print out the value of the variable “x”
	//ii. print out the type of the variable “x”

	// use CONVERSION to convert the TYPE of the VALUE stored in “x”
	//to the UNDERLYING TYPE
	//1. then use the “=” operator to ASSIGN that value to “y”
	//2. print out the value stored in “y”
	//3. print out the type of “y”

	fmt.Println(x)

	var y int
	y = int(x)

	fmt.Println(y)

	// print type of variables
	fmt.Println(x, reflect.TypeOf(x))
	fmt.Println(y, reflect.TypeOf(y))

	seperator()
}


func exercise3(){
	//1. Assign the following values to the three variables
	//a. for x assign 42
	//b. for y assign “James Bond”
	//c. for z assign true
	//a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
	//returned value of TYPE string using the short declaration operator to a
	//VARIABLE with the IDENTIFIER “s”
	//b. print out the value stored by variable “s”

	x:=42
	y:="James Bond"
	z:=true

	fmt.Sprintf("%v %v %v", x,y,z)

	seperator()
}



func exercise2(){
	// Use var to DECLARE three VARIABLES. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
	//variables and make sure the variables are of the following TYPE (meaning they can
	//store VALUES of that TYPE).
	//a. identifier “x” type int
	//b. identifier “y” type string
	//c. identifier “z” type bool
	//1. print out the values for each identifier
	//2. The compiler assigned values to the variables. What are these values called?

	var x int
	var y string
	var z bool

	// prints zero values for variables
	fmt.Println(x,y,z)

	seperator()
}

func exercise1(){
	// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS
	//a. identifier “x” type int
	//b. identifier “y” type string
	//c. identifier “z” type bool. And print the values stored in those variables using a. a single print statement b. multiple print statements.

	x := 55 	// Using the short declaration operator
	y := "hello!"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x,y,z) // Single line print
	seperator()
}


func seperator()  {
	fmt.Println("--------------------------------")
}