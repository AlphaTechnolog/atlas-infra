package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
)

type dynamoUrlAnalyticsTableRepo repository.PersistenceRepository[domain.ShortenedURL]

type URLListenerUseCase struct {
	dynamoUrlAnalyticsTable dynamoUrlAnalyticsTableRepo
}

func NewURLListenerUseCase(dynamoUrlAnalyticsTable dynamoUrlAnalyticsTableRepo) *URLListenerUseCase {
	return &URLListenerUseCase{
		dynamoUrlAnalyticsTable: dynamoUrlAnalyticsTable,
	}
}

func (u *URLListenerUseCase) HandleShortenedURL(ctx context.Context, entity *domain.ShortenedURL) error {
	if err := u.dynamoUrlAnalyticsTable.Create(ctx, *entity); err != nil {
		return fmt.Errorf("unable to insert entity into dynamodb table: %w", err)
	}

	rawEntity, _ := json.Marshal(entity)
	fmt.Println("saved successfully", string(rawEntity))

	return nil
}
