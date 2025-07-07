package handler

import (
	"context"
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
	_ = ctx

	for _, record := range e.Records {
		var shortenedURL domain.ShortenedURL
		if err := shortenedURL.FromSQSEvent(record); err != nil {
			return map[string]any{
				"error": err.Error(),
				"ok":    false,
			}, nil
		}

		if err := h.URLListenerUseCase.HandleShortenedURL(&shortenedURL); err != nil {
			return map[string]any{
				"error": err.Error(),
				"ok":    false,
			}, nil
		}
	}

	return map[string]any{"ok": true}, nil
}
