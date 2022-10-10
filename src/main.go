package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(req Request) (Response, error) {
	flags := make(map[string]bool)
	flags["POST_RATE_ADJUSTMENT_ACTION"] = true
	flags["POST_PRICE_RULE_ACTION"] = true
	return Response{Feature: req.Feature, Flag: flags[req.Feature]}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
