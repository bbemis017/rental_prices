package main

import (
	"fmt"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {

	fmt.Println("test")
	response, err := golambdainvoke.Run(golambdainvoke.Input{
		Port:    8001,
		Payload: "payload",
	})

	fmt.Println(err)
	fmt.Println(string(response))
}
