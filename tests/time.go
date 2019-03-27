package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	t = time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Println(time.Now().Year(), time.Now().Month())
}
