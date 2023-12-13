package main

import (
	"fmt"

	"goshaw/revstr"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(revstr.ReverseRunes("tac rekcah"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

// ref - https://go.dev/doc/code