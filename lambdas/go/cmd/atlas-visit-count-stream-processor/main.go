package main

import (
	"github.com/alphatechnolog/atlas-infra/cmd/atlas-visit-count-stream-processor/di"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := di.InitialiseVisitCountStreamProcessor()
	lambda.Start(h.Handle)
}
