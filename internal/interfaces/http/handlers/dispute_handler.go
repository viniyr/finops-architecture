package handlers

import (
	"encoding/json"
	"net/http"

	"aevyn/finops-arch/internal/usecases"
)

type OpenDisputeRequest struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Reason        string  `json:"reason"`
}

type DisputeHandler struct {
	Open *usecases.OpenDisputeUseCase
}

func NewDisputeHandler(open *usecases.OpenDisputeUseCase) DisputeHandler {
	return DisputeHandler{Open: open}
}

func (h DisputeHandler) OpenCardDispute(w http.ResponseWriter, r *http.Request) {
	var req OpenDisputeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Open.Execute(r.Context(), req.TransactionID, req.Amount, req.Reason); err != nil {
		http.Error(w, "failed to publish event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
