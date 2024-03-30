package main

import (
	"context"
	"fmt"

	"github.com/pravandkatyare/aws-sqs/aws"
)

const (
// queueURL = "https://sqs.ap-south-1.amazonaws.com/730335262455/test"
)

func main() {
	fmt.Println("Starting server")
	err := aws.NewSQSConnection().Connect().Subscribe(context.Background(), "test", nil)
	if err != nil {
		panic("error subscribing SQS")
	}
}
