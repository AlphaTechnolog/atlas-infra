package usecase

import (
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
)

type dynamoUrlAnalyticsTableRepo repository.PersistenceRepository[domain.ShortenedURL]

type URLConsumerUseCase struct {
	dynamoUrlAnalyticsTable dynamoUrlAnalyticsTableRepo
}

func NewURLConsumerUseCase(dynamoUrlAnalyticsTable dynamoUrlAnalyticsTableRepo) *URLConsumerUseCase {
	return &URLConsumerUseCase{
		dynamoUrlAnalyticsTable,
	}
}

func (u *URLConsumerUseCase) GetShortenedURL(urlID string) (domain.ShortenedURL, error) {
	var out domain.ShortenedURL
	if err := u.dynamoUrlAnalyticsTable.GetItemUnmarshal(urlID, &out); err != nil {
		return domain.ShortenedURL{}, fmt.Errorf("unable to obtain url by pk %s: %w", urlID, err)
	}

	return out, nil
}
