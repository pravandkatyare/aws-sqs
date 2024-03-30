package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSConnection struct {
	*Connection

	client *sqs.Client
}

func NewSQSConnection() *SQSConnection {
	return &SQSConnection{Connection: NewConnection()}
}
func (sqsc *SQSConnection) Connect() *SQSConnection {
	if sqsc.client == nil {
		sqsc.client = sqs.NewFromConfig(*sqsc.config)
	}
	return sqsc
}

func (sqsc *SQSConnection) Subscribe(ctx context.Context, queueName string, _ []string) error {
	queueURL, err := sqsc.client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: &queueName})
	if err != nil {
		return fmt.Errorf("unable to retrieve queue URL, err: %v", err)
	}
	fmt.Println("queueURL: ", *queueURL.QueueUrl)

	for {
		result, err := sqsc.client.ReceiveMessage(ctx,
			&sqs.ReceiveMessageInput{
				QueueUrl:              queueURL.QueueUrl,
				MessageAttributeNames: []string{"SentTimestamp"},
				AttributeNames:        []types.QueueAttributeName{types.QueueAttributeNameAll},
				WaitTimeSeconds:       0,
				MaxNumberOfMessages:   10,
				VisibilityTimeout:     0,
			})
		if err != nil {
			return fmt.Errorf("unable to receive message, err: %s", err)
		}

		for _, message := range result.Messages {
			fmt.Println("message: ", *message.Body)

		}
	}
}
