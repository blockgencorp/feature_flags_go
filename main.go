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
	flags := make(map[string]bool)
	flags["POST_RATE_ADJUSTMENT_ACTION"] = true
	flags["POST_PRICE_RULE_ACTION"] = true
	return Response{Feature: req.Feature, Flag: flags[req.Feature]}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
