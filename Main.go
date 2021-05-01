package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	fmt.Println("Logging the handler")
	return "Hello ƛ! whatsuppp", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
