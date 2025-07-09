package domain

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"time"
)

const DefaultTtlOffset = 30 * time.Minute

type ShortenedURL struct {
	ID         string `json:"id" dynamodbav:"id"`
	URL        string `json:"url" dynamodbav:"url"`
	VisitCount int64  `json:"visitCount" dynamodbav:"visitCount"`
	CreatedAt  string `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt  string `json:"updatedAt" dynamodbav:"updatedAt"`
	TTL        int64  `json:"-" dynamodbav:"ttl"`
}

func (s *ShortenedURL) FromSQSEvent(message events.SQSMessage) error {
	if err := json.Unmarshal([]byte(message.Body), s); err != nil {
		return fmt.Errorf("failed to parse shortened url from sqs queue: %w", err)
	}
	return nil
}

func (s *ShortenedURL) GenTTL() {
	s.TTL = time.Now().Local().Add(DefaultTtlOffset).Unix()
}

func (s *ShortenedURL) NeedsTTL() bool {
	return s.TTL == 0
}
