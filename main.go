package main

import (
	"fmt"
	"io/ioutil"

	_ "github.com/stretchr/testify"
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
