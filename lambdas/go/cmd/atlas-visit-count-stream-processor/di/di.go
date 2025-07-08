package di

import (
	"context"
	"github.com/alphatechnolog/atlas-infra/appconfig"
	apigateway "github.com/alphatechnolog/atlas-infra/infrastructure/apigateway"
	dynamodb "github.com/alphatechnolog/atlas-infra/infrastructure/dynamodb"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	"github.com/alphatechnolog/atlas-infra/internal/handler"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-visit-count-stream-processor"
	"log"
)

func InitialiseVisitCountStreamProcessor() *handler.VisitCountStreamProcessorHandler {
	dynamoClient, err := dynamodb.NewDynamoDBClient()
	if err != nil {
		log.Fatalf("failed to initialise dynamodb client: %v", err)
	}

	apiGatewayManagementAPIClient, err := apigateway.NewApiGatewayManagementAPI(appconfig.WebSocketEndpoint)
	if err != nil {
		log.Fatalf("failed to initialise apigateway client: %v", err)
	}

	wsConnsTableRepo := dynamodb.NewDynamoDBRepository[domain.WebSocketConnection](
		context.TODO(),
		dynamoClient,
		appconfig.WSConnectionsTableName,
		"connectionId",
	)

	u := usecase.NewVisitCountStreamProcessorUseCase(apiGatewayManagementAPIClient, dynamoClient, wsConnsTableRepo)

	return handler.NewVisitCountStreamProcessorHandler(u)
}
