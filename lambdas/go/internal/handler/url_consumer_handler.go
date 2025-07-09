package handler

import (
	"context"
	"fmt"
	"github.com/alphatechnolog/atlas-infra/pkg"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-consumer"
	"github.com/aws/aws-lambda-go/events"
)

type URLConsumerHandler struct {
	*usecase.URLConsumerUseCase
}

func NewURLConsumerHandler(urlConsumerUC *usecase.URLConsumerUseCase) *URLConsumerHandler {
	return &URLConsumerHandler{
		urlConsumerUC,
	}
}

type decodedRequest struct {
	ID string `json:"id"`
}

func (h *URLConsumerHandler) Handle(_ context.Context, e events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body decodedRequest
	if err := pkg.UnmarshalHTTPEventBody(e, &body); err != nil {
		return pkg.SendHTTPResponse(400, map[string]any{
			"error": fmt.Sprintf("cannot decode request body: %v", err),
			"ok":    false,
		}), nil
	}

	shortenedURL, err := h.URLConsumerUseCase.GetShortenedURL(body.ID)
	if err != nil {
		return pkg.SendHTTPResponse(404, map[string]any{
			"error": fmt.Sprintf("unable to obtain shortened url: %v", err),
			"ok":    false,
		}), nil
	}

	return pkg.SendHTTPResponse(200, map[string]any{
		"ok":           true,
		"shortenedUrl": shortenedURL,
	}), nil
}
