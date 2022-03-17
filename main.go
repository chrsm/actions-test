package main

import (
	"fmt"
	"io/ioutil"
)

// violate cuddle rule
func main() {
	byt, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Printf("panic-ing: %s", err)
		panic(err)
	}
	fmt.Println(byt)
}
