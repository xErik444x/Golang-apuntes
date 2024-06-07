package main

import (
	"flag"
	"fmt"
)

func main() {
	numero1 := flag.Int("numero1", 0, "primer numero")
	numero2 := flag.Int("numero2", 0, "primer numero")
	flag.Parse()
	res := *numero1 * *numero2
	fmt.Println(res)
}

// run go run .\main.go --numero1 5 --numero2 10
