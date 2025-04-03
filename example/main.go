package main

import (
	"fmt"

	"golang.org/x/net/http2"
)

func main() {
	fmt.Printf("%#v\n", http2.ClientConn{})
	fmt.Println("Hello, World!")
}
