package usecase

import (
	"encoding/json"
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
)

type URLListenerUseCase struct{}

func NewURLListenerUseCase() *URLListenerUseCase {
	return &URLListenerUseCase{}
}

func (u *URLListenerUseCase) HandleShortenedURL(event *domain.ShortenedURL) error {
	stringEvent, _ := json.Marshal(event)
	fmt.Println("got event", string(stringEvent))
	return nil
}
