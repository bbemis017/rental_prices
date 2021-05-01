package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

func hello() (string, error) {
	fmt.Println("Logging the handler")
	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}