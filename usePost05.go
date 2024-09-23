package main

import (
	"fmt"
	"github.com/carinfinin/pg05"
)

func main() {
	pg05.Hostname = "localhost"
	fmt.Println(pg05.Port)
	fmt.Println(pg05.Hostname)
}
