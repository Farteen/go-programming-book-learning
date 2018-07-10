package main

import (
	"fmt"
	"github.com/Farteen/Go-Programming-Book-Learning/2.5/tempconv"
)

func main() {
	var f tempconv.Fahrenheit = tempconv.Fahrenheit(1)
	var c tempconv.Celsius = tempconv.Celsius(2)

	fmt.Println(c.String())
	fmt.Println(f.String())
}