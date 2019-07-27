package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
    //"github.com/aws/aws-sdk-go/aws/credentials"
)
func main() {
	// use ~/.aws/credential file
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    
   /** use hard code credentail 
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String("us-west-2"),
        //Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
        Credentials: credentials.NewStaticCredentials("AKIAQVxxxxxxxxxxxxxx", "eSGe62tzxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",""),
    })*/
    
    svc := sqs.New(sess)

    // URL to our queue
    qURL := "https://sqs.us-west-2.amazonaws.com/USERNUMBER/THENAMEOFSQS"
    
    result, err := svc.SendMessage(&sqs.SendMessageInput{
    DelaySeconds: aws.Int64(10),
        MessageAttributes: map[string]*sqs.MessageAttributeValue{
            "Title": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("The Whistler"),
            },
            "Author": &sqs.MessageAttributeValue{
                DataType:    aws.String("String"),
                StringValue: aws.String("Chen Lin"),
            },
            "WeeksOn": &sqs.MessageAttributeValue{
                DataType:    aws.String("Number"),
                StringValue: aws.String("6"),
            },
        },
        MessageBody: aws.String("This is a testing message for Chen to test AWS SQS"),
        QueueUrl:    &qURL,
    })

    if err != nil {
        fmt.Println("Error", err)
        return
    }

    fmt.Println("Success", *result.MessageId)
}
