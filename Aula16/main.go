package main 

import (
	"fmt"
	"calculator/math"
)

func main() {
	sum := math.Sum(5, 6)
	sub := math.Sub(10, 4)

	fmt.Println(sum)
	fmt.Println(sub)
}