package main

import (
	// "unicode/utf8"
	// "fmt"
)

func main() {
	// const GoUsage = `...
	// this is
	// ...
	// `
	// fmt.Println(GoUsage)
	// strings.Contains(GoUsage, "this is")
	// s := "hello, 世界"
	// fmt.Println(len(s))
	// fmt.Println(utf8.RuneCountInString(s))

	// for i := 0; i < len(s); {
	// 	r, size := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%d\t%c\t%d\n",i, r, r)
	// 	i += size
	// }

	// for i, r := range "hello, 世界" {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r,r)
	// }


}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}