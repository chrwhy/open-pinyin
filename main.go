package main

import (
	"fmt"

	"github.com/chrwhy/open-pinyin/parser"
)

func main() {
	result := parser.Parse("oyhq")
	fmt.Println(result)
	initial := parser.ParseInitial("oyhq")
	fmt.Println(initial)
}
