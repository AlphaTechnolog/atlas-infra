package handler

import (
	"context"
	"github.com/alphatechnolog/atlas-infra/pkg"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-fetcher"
	"github.com/aws/aws-lambda-go/events"
)

type URLFetcherHandler struct {
	*usecase.URLFetcherUseCase
}

func NewURLFetcherHandler(useCase *usecase.URLFetcherUseCase) *URLFetcherHandler {
	return &URLFetcherHandler{useCase}
}

func (h *URLFetcherHandler) Handle(_ context.Context, _ events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	items, err := h.URLFetcherUseCase.GetItems()
	if err != nil {
		return pkg.SendHTTPResponse(400, map[string]any{
			"ok":  false,
			"err": err.Error(),
		}), nil
	}
	return pkg.SendHTTPResponse(200, map[string]any{
		"ok":    true,
		"items": items,
	}), nil
}
