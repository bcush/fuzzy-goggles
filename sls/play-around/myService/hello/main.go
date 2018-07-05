package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/net/context"
)

func Handler() {
	fmt.Println(context.Context)
	return
}

func main() {
	lambda.Start(Handler)
}
