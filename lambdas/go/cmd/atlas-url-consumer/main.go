package main

import (
	"github.com/alphatechnolog/atlas-infra/cmd/atlas-url-consumer/di"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := di.InitialiseAtlasURLConsumerDependencies()
	lambda.Start(h.Handle)
}
