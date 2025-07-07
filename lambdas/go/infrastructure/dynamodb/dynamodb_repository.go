package infrastructure

import (
	"context"
	"errors"
	"fmt"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBRepository[T any] struct {
	client     *dynamodb.Client
	tableName  string
	primaryKey string
}

func NewDynamoDBRepository[T any](
	client *dynamodb.Client,
	tableName string,
	primaryKey string,
) repository.PersistenceRepository[T] {
	return &DynamoDBRepository[T]{
		tableName:  tableName,
		primaryKey: primaryKey,
		client:     client,
	}
}

func (r *DynamoDBRepository[T]) GetByPK(pk string) (map[string]types.AttributeValue, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			r.primaryKey: &types.AttributeValueMemberS{Value: pk},
		},
	}

	result, err := r.client.GetItem(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item by %s: %w", pk, err)
	}

	if len(result.Item) == 0 {
		return nil, fmt.Errorf("failed to get item by %s: len 0", pk)
	}

	return result.Item, nil
}

func (r *DynamoDBRepository[T]) GetItemUnmarshal(pk string, out any) error {
	item, err := r.GetByPK(pk)
	if err != nil {
		return err
	}

	return attributevalue.UnmarshalMap(item, out)
}

func (r *DynamoDBRepository[T]) Create(ctx context.Context, entity T) error {
	item, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName:           &r.tableName,
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	if _, err := r.client.PutItem(ctx, input); err != nil {
		var condFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &condFailedErr) {
			return fmt.Errorf("item with the same id already exists")
		}
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}
