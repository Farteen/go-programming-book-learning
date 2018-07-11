package main

import (
	"os"
	// "fmt"
	// "math"
	// "regexp"
	"runtime"
)

func main() {
	// switch s := suit(drawCard()); s {
	// case "Spades":
	// case "Hearts":
	// case "Diamonds":
	// case "Clubs":
	// default:
	// 	panic(fmt.Sprintf("invalid suit %q", s))
	// }
	// var httpScheme = regexp.MustCompile(`https?:`)
	// result := httpScheme.FindAllStringIndex("http://www.baidu.com https://www.baidu.com htt://www.baidu.com httttttt:///www.baidu.com this is asfjsalkf http://333.rrr this", math.MaxUint32)
	// fmt.Println(result)

		defer printStack() 
		// f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// func (p Point) X() int {
// 	return p.X
// }