package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PersistenceRepository[T any] interface {
	// GetByPK FIXME: We should not be using dynamo types here, this should be truly generic.
	GetByPK(pk string) (map[string]types.AttributeValue, error)
	GetItemUnmarshal(pk string, out any) error
	Create(entity T) error
	UpdateAndGet(pk string, setters []string, variables map[string]any, out *T) error
}
