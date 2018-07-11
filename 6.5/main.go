package main

func main() {

}

type Inset struct {
	words []int64
}

func (*Inset) Count() int64 {
	return 1
}