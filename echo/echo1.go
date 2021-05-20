package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	fmt.Println(os.Args[0])
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("%d %s\n", i+1, arg)
	}
	fmt.Println(s)
}
