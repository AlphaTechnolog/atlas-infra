package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PersistenceRepository[T any] interface {
	GetByPK(pk string) (map[string]types.AttributeValue, error)
	GetItemUnmarshal(pk string, out any) error
	Create(ctx context.Context, entity T) error
}
