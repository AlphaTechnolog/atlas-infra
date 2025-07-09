package main

import (
	"github.com/alphatechnolog/atlas-infra/cmd/atlas-url-fetcher/di"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := di.InitialiseAtlasURLFetcherDependencies()
	lambda.Start(h.Handle)
}
