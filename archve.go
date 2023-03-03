package main

import (
	"flag"
	"fmt"
)

func main2() {
	flag.Parse()
	arg := flag.Arg(0)
	fmt.Printf("Hello %s\n", arg)
}
