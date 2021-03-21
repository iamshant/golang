package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(Multiple3And5(10))
	// fmt.Println(GetCount("string"))
	// a := "103 123 4444 99 2000"
	// fmt.Println(Closest(a))
	// fmt.Println(DigitalRoot(1))
	fmt.Println(FindNb(1071225))
}

// func Multiple3And5(number int) int {
// 	i := 0
// 	sum := 0
// 	for i = 2; i < number; i++ {
// 		if i%3 == 0 || i%5 == 0 {
// 			sum = sum + i
// 		}
// 		// fmt.Println(i)
// 	}
// 	return sum
// }
// func Multiple3And5(number int) int {
// 	n3 := (number - 1) / 3
// 	n5 := (number - 1) / 5
// 	n15 := (number - 1) / 15
// 	return (3*n3*(n3+1) + 5*n5*(n5+1) - 15*n15*(n15+1)) / 2
// }
// func GetCount(str string) (count int) {
// 	// Enter solution here
// 	i := 0
// 	for i = 0; i < len(str); i++ {
// 		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
// 			count++
// 		}
// 	}

// 	return count
// }

// str = "103 123 4444 99 2000"
// func Closest(str string) string {
// 	// your code
// 	fmt.Println(len(str))

// 	s := strings.Split(str, " ")
// 	fmt.Println(s)
// 	// i := 0
// 	for i := 0; i < len(s); i++ {
// 		fmt.Println(s[i])
// 	}
// 	return str
// }

// func DigitalRoot(n int) int {
// 	str := strconv.Itoa(a)""
// 	k := 0
// 	if len(str) > 1 {
// 		for i := 0; i < len(str); i++ {
// 			k = int(str[i]) + k
// 		}
// 	}
// 	return k
// }

func FindNb(m int) int {
	// your code
	i := 1
	// fmt.Println(5 - 5.4)
	// a := int(math.Pow(float64(i), 3))
	// fmt.Println(a)
	for {
		m = m - int(math.Pow(float64(i), (3)))
		if m < 0 {
			return -1
		} else if m == 0 {
			return i
		}
		i++
	}
}
