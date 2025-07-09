package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	ddbconversions "github.com/aereal/go-dynamodb-attribute-conversions/v2"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type dynamoWsConnsTableRepo repository.PersistenceRepository[domain.WebSocketConnection]

type VisitCountStreamProcessorUseCase struct {
	apiGatewayManagementAPIClient *apigatewaymanagementapi.Client
	dynamoClient                  *dynamodb.Client
	wsConnsTableRepo              dynamoWsConnsTableRepo
}

func NewVisitCountStreamProcessorUseCase(
	apiGatewayManagementAPIClient *apigatewaymanagementapi.Client,
	dynamoClient *dynamodb.Client,
	wsConnsTableRepo dynamoWsConnsTableRepo,
) *VisitCountStreamProcessorUseCase {
	return &VisitCountStreamProcessorUseCase{
		apiGatewayManagementAPIClient: apiGatewayManagementAPIClient,
		dynamoClient:                  dynamoClient,
		wsConnsTableRepo:              wsConnsTableRepo,
	}
}

func (u *VisitCountStreamProcessorUseCase) getConnections() ([]domain.WebSocketConnection, error) {
	var connections []domain.WebSocketConnection
	if err := u.wsConnsTableRepo.GetItems(&connections); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dynamodb items: %w", err)
	}
	return connections, nil
}

func (u *VisitCountStreamProcessorUseCase) Handle(ctx context.Context, updateRecord events.DynamoDBEventRecord) error {
	_ = ctx

	log.Printf("Processing record ID: %s, eventName: %s\n", updateRecord.EventID, updateRecord.EventName)

	// dynamo sdk needs types updates i guess 😭
	convNew := ddbconversions.AttributeValueMapFrom(updateRecord.Change.NewImage)
	convOld := ddbconversions.AttributeValueMapFrom(updateRecord.Change.OldImage)
	var newRecord domain.ShortenedURL
	var oldRecord domain.ShortenedURL
	if err := attributevalue.UnmarshalMap(convNew, &newRecord); err != nil {
		return fmt.Errorf("failed to unmarshal new image: %w", err)
	}
	if err := attributevalue.UnmarshalMap(convOld, &oldRecord); err != nil {
		return fmt.Errorf("failed to unmarshal old image: %w", err)
	}

	if newRecord.VisitCount == oldRecord.VisitCount {
		log.Printf("visit count for %s did not change. Skipping", newRecord.URL)
		return nil
	}

	connections, err := u.getConnections()
	if err != nil {
		return fmt.Errorf("failed to get websocket active connections: %w", err)
	}

	message := map[string]any{
		"action":     "visitCountUpdate",
		"urlId":      newRecord.ID,
		"visitCount": newRecord.VisitCount,
	}
	messageBytes, _ := json.Marshal(message)

	for _, conn := range connections {
		log.Printf("Broadcasting update to connection ID: %s", conn.ConnectionID)
		_, err = u.apiGatewayManagementAPIClient.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(conn.ConnectionID),
			Data:         messageBytes,
		})
		if err != nil {
			log.Printf("failed to send message to connection ID: %s, error: %v", conn.ConnectionID, err)
		}
	}

	return nil
}
