package main 

import (
	"fmt"
	"calculator/math"
	"github.com/labstack/echo/v4"
)

func main() {
	// echo instace
	e := echo.New()
	e.Listener.Close()

	sum := math.Sum(5, 6)
	sub := math.Sub(10, 4)

	fmt.Println(sum)
	fmt.Println(sub)
}