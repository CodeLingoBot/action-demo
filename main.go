package main

import (
	"context"
	"fmt"
)

func main() {
	Thing()
	foo("Hello", context.Background())
}

func Thing() {
	t := []string{}
	fmt.Println(t)
}
//blah
func foo(msg string, ctx context.Context) {
	fmt.Println(msg)
}
