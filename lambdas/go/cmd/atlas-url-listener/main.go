package main

import (
	"github.com/alphatechnolog/atlas-infra/cmd/atlas-url-listener/di"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := di.InitialiseAtlasURLListenerDependencies()
	lambda.Start(h.Handle)
}
