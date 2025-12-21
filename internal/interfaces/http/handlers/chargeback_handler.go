package handlers

import (
	"encoding/json"
	"net/http"

	"aevyn/finops-arch/internal/usecases"
)

type OpenChargebackRequest struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Reason        string  `json:"reason"`
}

func OpenChargeback(w http.ResponseWriter, r *http.Request) {
	var req OpenChargebackRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := usecases.OpenChargeback(r.Context(), req.TransactionID, req.Amount, req.Reason)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
