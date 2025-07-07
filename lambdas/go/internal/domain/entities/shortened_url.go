package domain

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type ShortenedURL struct {
	ID        string `json:"id" dynamodbav:"id"`
	URL       string `json:"url" dynamodbav:"url"`
	CreatedAt string `json:"createdAt" dynamodbav:"createdAt"`
}

func (s *ShortenedURL) FromSQSEvent(message events.SQSMessage) error {
	if err := json.Unmarshal([]byte(message.Body), s); err != nil {
		return fmt.Errorf("failed to parse shortened url from sqs queue: %w", err)
	}
	return nil
}
