package usecase

import (
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
)

type dynamoURLAnalyticsTableRepo repository.PersistenceRepository[domain.ShortenedURL]

type URLFetcherUseCase struct {
	dynamoURLAnalyticsTableRepo
}

func NewURLFetcherUseCase(dynamoURLAnalyticsTableRepo dynamoURLAnalyticsTableRepo) *URLFetcherUseCase {
	return &URLFetcherUseCase{
		dynamoURLAnalyticsTableRepo: dynamoURLAnalyticsTableRepo,
	}
}

func (u *URLFetcherUseCase) GetItems() ([]domain.ShortenedURL, error) {
	var urls []domain.ShortenedURL
	if err := u.dynamoURLAnalyticsTableRepo.GetItems(&urls); err != nil {
		return nil, fmt.Errorf("cannot obtain all urls from dynamo db: %w", err)
	}
	return urls, nil
}
