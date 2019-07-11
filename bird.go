package main

import "fmt"

func main(){
	b := Bird{}
	b.colour = "white"
	b.weight = 30
	b.Tail.length = 20
	b.Tail.width = 4
	b.Weng.length = 10
	b.Weng.square = 100

	fmt.Println(b)

}

type (
	Weng struct {
		square float32
		length float32
	}

	Tail struct {
		width  float32
		length float32
	}
)

type (
	Bird struct {
		Tail   Tail
		Weng   Weng
		colour string
		weight float32
	}
)