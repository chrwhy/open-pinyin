package main

import (
	"github.com/chrwhy/open-pinyin/parser"
	"log"
)

func Test(input string) {
	result := parser.Parse(input)
	log.Println(input, ":")
	for i, _ := range result {
		log.Println(result[i])
	}
}

func main() {
	Test("ceshi")
}
