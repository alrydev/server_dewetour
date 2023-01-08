package transactionsdto

import "dewe/models"

type TransactionResponse struct {
	ID         int         `json:"id"`
	CounterQty int         `json:"counter_qty"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	Attachment string      `json:"image"`
	Trip       models.Trip `json:"tripId"`
}
