package classes

type Action struct {
	ID        int64  `json:"id"`
	DateTime  int64  `json:"date_time"`
	IP        string `json:"ip"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	SessionID string `json:"session_id"`
}
