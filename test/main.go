package main

import (
	"fmt"
	"goreloaded"
	"strings"
)

func main() {
	s := strings.Split("I have an banana, a apple, and a orange.", " ")
	fmt.Println(strings.Join(goreloaded.CheckAorAn(s), " "))
}