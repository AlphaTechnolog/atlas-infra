package di

import (
	"context"
	"github.com/alphatechnolog/atlas-infra/appconfig"
	infrastructure "github.com/alphatechnolog/atlas-infra/infrastructure/dynamodb"
	"log"

	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	"github.com/alphatechnolog/atlas-infra/internal/handler"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-listener"
)

func InitialiseAtlasURLListenerDependencies() *handler.URLListenerHandler {
	dynamoClient, err := infrastructure.NewDynamoDBClient()
	if err != nil {
		log.Fatalf("failed to initialiase dynamodb client: %v", err)
	}

	urlAnalyticsTable := infrastructure.NewDynamoDBRepository[domain.ShortenedURL](
		context.TODO(),
		dynamoClient,
		appconfig.DynamoUrlAnalyticsTableName,
		"id",
	)

	uc := usecase.NewURLListenerUseCase(urlAnalyticsTable)

	return handler.NewURLListenerHandler(uc)
}
