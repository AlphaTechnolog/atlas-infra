package usecase

import (
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	repository "github.com/alphatechnolog/atlas-infra/internal/repository/persistence"
	"time"
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
	setters := []string{"visitCount = visitCount + :inc", "updatedAt = :now"}
	attrValues := map[string]any{
		":inc": 1,
		":now": time.Now().UTC().Format(time.RFC3339),
	}

	var out domain.ShortenedURL
	if err := u.dynamoUrlAnalyticsTable.UpdateAndGet(urlID, setters, attrValues, &out); err != nil {
		return domain.ShortenedURL{}, fmt.Errorf("unable to obtain url by pk: %s: %w", urlID, err)
	}

	return out, nil
}
