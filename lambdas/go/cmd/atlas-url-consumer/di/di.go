package di

import (
	"github.com/alphatechnolog/atlas-infra/appconfig"
	infrastructure "github.com/alphatechnolog/atlas-infra/infrastructure/dynamodb"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	"github.com/alphatechnolog/atlas-infra/internal/handler"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-consumer"
	"log"
)

func InitialiseAtlasURLConsumerDependencies() *handler.URLConsumerHandler {
	dynamoClient, err := infrastructure.NewDynamoDBClient()
	if err != nil {
		log.Fatalf("failed to initialise dynamodb client: %v", err)
	}

	urlAnalyticsTable := infrastructure.NewDynamoDBRepository[domain.ShortenedURL](
		dynamoClient,
		appconfig.DynamoUrlAnalyticsTableName,
		"id",
	)

	uc := usecase.NewURLConsumerUseCase(urlAnalyticsTable)

	return handler.NewURLConsumerHandler(uc)
}
