package main

import (
	"strings"
	// "strings"
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
)

/// this compiles with error 
/// name := "this is outside the func"

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n",n, line)
			}
		}
	}


	// counts := make(map[string]int)
	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// 	for line, n := range counts {
	// 		if n > 1 {
	// 			fmt.Printf("%d\t%s\n",n, line)
	// 		}
	// 	}
	// }
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }	
	// fmt.Println(s)
	// fmt.Println(name)

	// fmt.Println(strings.Join(os.Args[1:], " "))

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)

	// for input.Scan() {
	// 	counts[input.Text()]++
	// }

	// for line, n := range counts {
	// 	if n > 0 {
	// 		fmt.Println("%d\t%s\n",n, line)
	// 	}
	// }
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	if input.Scan() {
		counts[input.Text()]++
	}
}