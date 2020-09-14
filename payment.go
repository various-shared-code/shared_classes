package classes

type Payment struct {
	ID        int64  `json:"id"`
	DateTime  int64  `json:"date_time"`
	IP        string `json:"ip"`
	UserID    int64  `json:"user_id"`
	Amount    int64  `json:"amount"`
	Currency  int64  `json:"currency"`
	AmountUSD int64  `json:"amount_usd"`
}
