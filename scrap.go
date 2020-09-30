package main

import (
	"fmt"

)

func functionName(param1 string, param2 int) {}

func functionName1(param1, param2 int) {}

func functionName2() int {
	return 42
}

func retunMulti() (int, string) {
	return 42, "foobar"
}

var x, str = retunMulti()

func retunMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	return
}

var i, z = retunMulti2()

func addboth() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 8))
}

func scope() func() int {
	outerVar := 2
	foo := func() int {
		return outerVar
	}
	return foo
}

// func anotherScope() func() int {
// 	outerVar = 444
// 	return foo
// }

// var foo int
// var foo int = 24
// var foo, bar int = 32, 545
// var foo = 23
// foo := 45
// const constant = "this is permanent."

const (
	_ = iota
	a
	b
	c = 1 << iota
	d
)

func adder(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// num := 42
// f :=  float64(num)
// u := uint(f)
var num int = 42
var f float64 = float64(num)
var u uint = uint(f)

// fmt.Printf("%T,%T,%T", num, f, u)
// for i := 1; i < 10; i++ {
//     fmt.Println(i)
// }

// for ; i < 10; {
// }

// for i < 10 {
// }

// if true {
// 	for i := 0; i < 2; i++ {
// 		for j := i; j < 3; j++ {
// 			if i == 0 {
// 				continue
// 			}
// 			fmt.Println(j)
// 			if j == 2 {
// 				break
// 			}
// 		}
// 	}
// }

func main() {
	var x [10]int
	x[3] = 42
	i := x[3]

	// var a1 = [2]int{1, 2}
	// a1 := [2]int{1,2}
	a1 := [...]int{1,2,3,4,5,6,7,8,9}

	fmt.Println(x)
	fmt.Println(i)
	fmt.Println(a1)
}
