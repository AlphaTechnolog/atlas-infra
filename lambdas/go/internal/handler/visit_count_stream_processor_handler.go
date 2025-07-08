package handler

import (
	"context"
	"fmt"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-visit-count-stream-processor"
	"github.com/aws/aws-lambda-go/events"
)

type VisitCountStreamProcessorHandler struct {
	*usecase.VisitCountStreamProcessorUseCase
}

func NewVisitCountStreamProcessorHandler(
	visitCountStreamProcessorUC *usecase.VisitCountStreamProcessorUseCase,
) *VisitCountStreamProcessorHandler {
	return &VisitCountStreamProcessorHandler{
		visitCountStreamProcessorUC,
	}
}

func (h *VisitCountStreamProcessorHandler) Handle(ctx context.Context, e events.DynamoDBEvent) {
	for _, record := range e.Records {
		if record.EventName == "MODIFY" {
			if err := h.VisitCountStreamProcessorUseCase.Handle(ctx, record); err != nil {
				fmt.Printf("Unable to handle updated record: %s (skipping): %v", record.EventID, err)
				continue
			}
		}
	}
}
