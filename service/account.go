package service

type NewAccountRequest struct {
	AmountType string  `json:"amount_type"`
	Amount     float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID   int     `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
	AmountType  string  `json:"amount_type"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccount(int) ([]AccountResponse, error)
}
