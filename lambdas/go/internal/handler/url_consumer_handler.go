package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	URL string `json:"url"`
}

func decodeRequest(request events.APIGatewayV2HTTPRequest) (decodedRequest, error) {
	var decoded decodedRequest
	requestBody := []byte(request.Body)
	if request.IsBase64Encoded {
		data, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			return decodedRequest{}, fmt.Errorf("failed to decode b64 from api gateway request: %w", err)
		}
		requestBody = data
	}
	if err := json.Unmarshal(requestBody, &decoded); err != nil {
		return decodedRequest{}, fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	return decoded, nil
}

func sendResponse[T any](status int, body T) events.APIGatewayV2HTTPResponse {
	compressedBody, _ := json.Marshal(body)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       string(compressedBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (h *URLConsumerHandler) Handle(ctx context.Context, e events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	_ = ctx

	body, err := decodeRequest(e)
	if err != nil {
		return sendResponse(400, map[string]any{
			"error": fmt.Sprintf("cannot decode request body: %v", err),
			"ok":    false,
		}), nil
	}

	// TODO: Increment shortenedURL.VisitCount++
	shortenedURL, err := h.URLConsumerUseCase.GetShortenedURL(body.URL)
	if err != nil {
		return sendResponse(404, map[string]any{
			"error": fmt.Sprintf("unable to obtain shortened url: %v", err),
			"ok":    false,
		}), nil
	}

	return sendResponse(200, map[string]any{
		"ok":           true,
		"shortenedUrl": shortenedURL,
	}), nil
}
