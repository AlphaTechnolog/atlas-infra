package di

import (
	"context"
	"github.com/alphatechnolog/atlas-infra/appconfig"
	infrastructure "github.com/alphatechnolog/atlas-infra/infrastructure/dynamodb"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	"github.com/alphatechnolog/atlas-infra/internal/handler"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-fetcher"
	"log"
)

func InitialiseAtlasURLFetcherDependencies() *handler.URLFetcherHandler {
	dynamoClient, err := infrastructure.NewDynamoDBClient()
	if err != nil {
		log.Fatalf("unable to create new dynamo client: %v", err)
	}

	dynamoURLAnalyticsTableRepo := infrastructure.NewDynamoDBRepository[domain.ShortenedURL](
		context.TODO(),
		dynamoClient,
		appconfig.DynamoUrlAnalyticsTableName,
		"id",
	)

	useCase := usecase.NewURLFetcherUseCase(dynamoURLAnalyticsTableRepo)

	return handler.NewURLFetcherHandler(useCase)
}
