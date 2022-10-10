package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Feature string `json:"feature"`
}

type Response struct {
	Feature string `json:"feature"`
	Flag    bool   `json:"flag"`
}

func HandleLambdaEvent(req Request) (Response, error) {
	return Response{Feature: req.Feature, Flag: false}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
