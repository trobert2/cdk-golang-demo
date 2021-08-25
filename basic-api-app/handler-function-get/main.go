package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

const version int64 = 1

type MyEvent struct {
	SSMParameterPath string `json:"ssm-parameter-path"`
	Value            string `json:"value"`
}

type Response struct {
	Status  string `json:"status"`
	Version int64  `json:"version"`
	MyEvent
}

func init() {
	//	awsSess, err := session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable})
	//	if err != nil {
	//		log.Fatal(errors.New("failed to open aws session"))
	//	}

	// XRay Setup
	// xray.AWS(ec2Client.Client)
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {

	finalResponse := Response{
		Status:  "success",
		Version: version,
		MyEvent: MyEvent{
			SSMParameterPath: event.SSMParameterPath,
			Value:            event.Value,
		},
	}

	finalResponseBytes, err := json.Marshal(finalResponse)
	if err != nil {
		log.Fatal(err)
	}

	return string(finalResponseBytes), nil
}

func main() {
	lambda.Start(HandleRequest)
}
