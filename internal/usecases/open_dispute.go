package usecases

import (
	"context"
	"encoding/json"
	"time"

	"aevyn/finops-arch/internal/interfaces/messaging"
)

type OpenDisputeUseCase struct {
	Pub messaging.Publisher
}

type DisputeOpenedV1 struct {
	EventID       string    `json:"event_id"`
	DisputeID     string    `json:"dispute_id"`
	Rail          string    `json:"rail"`
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Reason        string    `json:"reason"`
	OccurredAt    time.Time `json:"occurred_at"`
}

func (uc *OpenDisputeUseCase) Execute(ctx context.Context, txID string, amount float64, reason string) error {
	ev := DisputeOpenedV1{
		EventID:       "todo-uuid",
		DisputeID:     "todo-id",
		Rail:          "CARD",
		TransactionID: txID,
		Amount:        amount,
		Reason:        reason,
		OccurredAt:    time.Now().UTC(),
	}

	b, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	return uc.Pub.Publish(ctx, "dispute.opened.card.v1", b)
}
