package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type ShortenedURL struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *ShortenedURL) FromSQSEvent(message events.SQSMessage) error {
	if err := json.Unmarshal([]byte(message.Body), s); err != nil {
		return fmt.Errorf("failed to parse shortened url from sqs queue: %w", err)
	}
	return nil
}
