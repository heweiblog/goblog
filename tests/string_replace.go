package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hhhh\r\njkhshhjd\r\nshjhjfjf\r\n"
	fmt.Println(strings.Replace(s, "\r\n", "<br>", -1))
}
