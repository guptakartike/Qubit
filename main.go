package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	n:="abc@gmail.com"
	fmt.Println(len(n))
	fmt.Println(utf8.RuneCountInString(n))
}