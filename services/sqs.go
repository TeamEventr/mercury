package services

//
// import (
// 	"fmt"
//
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/sqs"
// )
//
// // sqsClient := sqs.New()
//
// func sendMessage(svc *sqs.SQS, queueUrl string, body string) {
// 	sendMsgInput := &sqs.SendMessageInput{
// 		QueueUrl:    aws.String(queueUrl),
// 		MessageBody: aws.String(body),
// 	}
//
// 	result, err := svc.SendMessage(sendMsgInput)
// 	if err != nil {
// 		// Setup better logging
// 		fmt.Println("Error", err)
// 		return
// 	}
// 	fmt.Print("Message ID: %s\n", *result.MessageId)
// }
