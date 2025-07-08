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
	"reflect"
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

func (r *DynamoDBRepository[T]) Create(entity T) error {
	item, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName:           &r.tableName,
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	if _, err := r.client.PutItem(context.TODO(), input); err != nil {
		var condFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &condFailedErr) {
			return fmt.Errorf("item with the same id already exists")
		}
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (r *DynamoDBRepository[T]) UpdateAndGet(pk string, setters []string, variables map[string]any, out *T) error {
	updateExpression := "SET "
	for i, expr := range setters {
		if i == 0 {
			updateExpression += expr
		} else {
			updateExpression += fmt.Sprintf(", %s", expr)
		}
	}

	attrValues := map[string]types.AttributeValue{}
	for key, variable := range variables {
		switch variable.(type) {
		case string:
			attrValues[key] = &types.AttributeValueMemberS{
				Value: variable.(string),
			}
		case int, int64:
			attrValues[key] = &types.AttributeValueMemberN{
				Value: fmt.Sprint(variable.(int)),
			}
		case bool:
			attrValues[key] = &types.AttributeValueMemberBOOL{
				Value: variable.(bool),
			}
		case float32, float64:
			attrValues[key] = &types.AttributeValueMemberN{
				Value: fmt.Sprint(variable.(float64)),
			}
		default:
			return fmt.Errorf("attribute type not implemented for attrValues: %s", reflect.TypeOf(variable))
		}
	}

	key, err := attributevalue.MarshalMap(map[string]string{r.primaryKey: pk})
	if err != nil {
		return fmt.Errorf("failed to marshal key %s:%s: %w", r.primaryKey, pk, err)
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(r.tableName),
		Key:                       key,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: attrValues,
		ReturnValues:              "ALL_NEW",
	}

	result, err := r.client.UpdateItem(context.TODO(), updateInput)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	if result.Attributes == nil {
		return fmt.Errorf("failed to update item: no attributes found")
	}

	err = attributevalue.UnmarshalMap(result.Attributes, out)
	if err != nil {
		return fmt.Errorf("failed to serialise dynamo response: %w", err)
	}

	return nil
}
