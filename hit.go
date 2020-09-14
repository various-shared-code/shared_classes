package classes

import "encoding/json"

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

func (p Hit) ToJson() []byte {
	js, _ := json.Marshal(p)
	return js
}

func (p Hit) ToJsonString() string {
	return string(p.ToJson())
}
