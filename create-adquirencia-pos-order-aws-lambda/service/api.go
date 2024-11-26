package service

type (
	CreateOrderRequest struct {
		OrderID   string  `json:"order_id"`
		AccountID string  `json:"account_id"`
		Currency  string  `json:"currency"`
		Amount    float64 `json:"amount"`
		RefNumber string  `json:"ref_number"`
		Taxes     []*Tax  `json:"taxes"`
		Origin    string  `json:"origin"`
		CollabID  string  `json:"collab_id"`
	}

	Tax struct {
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}
)
