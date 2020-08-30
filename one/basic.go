package main

import (
	"fmt"
	"math"

)

func mathFunc() {

	for x := 0; x < 80; x++ {
		fmt.Printf("x = %d, e = %8.30f\n", x, math.Exp(float64(x)))
	}

}
