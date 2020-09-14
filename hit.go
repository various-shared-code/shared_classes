package classes

type Hit struct {
	ID         int64  `json:"id"`
	DateTime   int64  `json:"date_time"`
	IP         string `json:"ip"`
	Domain     string `json:"domain"`
	CampaignID int64  `json:"campaign_id"`
	PageID     int64  `json:"page_id"`
	ServerHit  bool   `json:"server_hit"`
	SessionID  string `json:"session_id"`
}
