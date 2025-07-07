package handler

import (
	"context"
	"fmt"
	domain "github.com/alphatechnolog/atlas-infra/internal/domain/entities"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-listener"
	"github.com/aws/aws-lambda-go/events"
)

type URLListenerHandler struct {
	*usecase.URLListenerUseCase
}

func NewURLListenerHandler(urlListenerUC *usecase.URLListenerUseCase) *URLListenerHandler {
	return &URLListenerHandler{
		urlListenerUC,
	}
}

func (h *URLListenerHandler) Handle(ctx context.Context, e events.SQSEvent) (map[string]any, error) {
	for _, record := range e.Records {
		var shortenedURL domain.ShortenedURL
		if err := shortenedURL.FromSQSEvent(record); err != nil {
			return nil, fmt.Errorf("unable to parse shortened url from sqs: %w", err)
		}

		if err := h.URLListenerUseCase.HandleShortenedURL(ctx, &shortenedURL); err != nil {
			return nil, fmt.Errorf("unable to handle shortened url: %w", err)
		}
	}

	return map[string]any{"ok": true}, nil
}
