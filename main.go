package main

import (
	"fmt"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

//interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Parameter() float64 {
	return 2 * (r.width + r.height)
}
func main() {

	//interface
	var sh Shape
	// fmt.Println("Value of s is", sh)
	// fmt.Printf("type of S is %T\n", sh)

	sh = Rect{5.0, 4.0}
	rh := Rect{5.0, 4.0}

	fmt.Printf("type of s is %T\n", sh)
	fmt.Printf("value of s is %v\n", sh)
	fmt.Println("area of rectangle s", sh.Area())
	fmt.Println("s==r is", sh == rh)
	// fmt.Println("Hello Shant")

	// var age int64 = 18

	// var string1 string = "My name is Shant"
	// var first_int, second_int int = 10, 20
	// var first_float float64 = 10.2

	// var int1 uint8 = 100
	// var int2 int8 = -100
	// var x string = fmt.Sprintf("Hello Golang  %v %t\n", 10, true)

	// // Taking input from user
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()

	// fmt.Println("Type the year you born: ")
	// scanner.Scan()
	// year, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	// fmt.Printf("you typed:   %q\n\n\n", input)
	// fmt.Printf("\nYou will be %d years old at the end of 2021\n", 2021-year)

	// // condition: if, else if, else
	// fmt.Println("if else")
	// if 2021-year > age {
	// 	fmt.Println("You can ride\n")
	// } else if (2021-year) < 18 && (2021-year) > 14 {
	// 	fmt.Println("You can ride with your parents\n")
	// } else {
	// 	fmt.Println("You can't ride alone\n")
	// }

	// fmt.Println("for loop")
	// // for loop
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println("\n")

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println("\n")

	// fmt.Println("switch version 1")
	// // switch
	// ans := -2

	// switch ans {
	// case 1, -1, -2:
	// 	fmt.Println("1. One")
	// 	fmt.Println("1. One\n")
	// case 2:
	// 	fmt.Println("2. Two\n")
	// case 3:
	// 	fmt.Println("Shant\n")
	// default:
	// 	fmt.Println("no case match\n")
	// }

	// fmt.Println("switch version 2")
	// switch {
	// case ans == -1:
	// 	fmt.Println("1. One")
	// 	fmt.Println("1. One\n")
	// case ans == 2:
	// 	fmt.Println("2. Two\n")
	// case ans == 3:
	// 	fmt.Println("Shant\n")
	// default:
	// 	fmt.Println("no case match\n")
	// }

	// // array[]
	// fmt.Println("array")
	// var arr [5]int
	// arr[4] = 100

	// arr1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// fmt.Println(arr1)
	// fmt.Printf("length of the arr is %d\n", len(arr))

	// // we can't append an array, but we can do sth like below in a sliced array
	// var b []int = []int{10, 20, 30}
	// b = append(b, 9999)
	// fmt.Println(b)

	// // 2D array
	// arr2D := [2][3]int{{1, 2, 1}, {3, 4, 5}}
	// fmt.Println(arr2D)

	// // Slicing
	// var s []int = arr1[1:3]
	// fmt.Println(s)
	// //cap -> capacity
	// fmt.Println(s[:cap(s)])
	// fmt.Println("\n")

	// // struct
	// fmt.Println("Struct")
	// fmt.Println(rectangle{10.2, 5.3, "red"})

	// fmt.Printf("%v\n", 2511110)
	// fmt.Printf(x)
	// fmt.Println(int1)
	// fmt.Println(int2)
	// fmt.Println(first_int+second_int, first_float)
	// fmt.Println(string1)

	// var w1 int
	// w1 = 200000000000
	// fmt.Println(w1)

	//interface
	var sh Shape
	fmt.Println("Value of s is", sh)
	fmt.Printf("type of S is %T\n", sh)

	sh = Rect{5.0, 4.0}
	rh := Rect{5.0, 4.0}

	fmt.Printf("type of s is %T\n", sh)
	fmt.Printf("value of s is %v\n", sh)
	fmt.Println("area of rectangle s", sh.Area())
	fmt.Println("s==r is", sh == rh)
}
