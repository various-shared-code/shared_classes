package classes

import "encoding/json"

type Domain struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	UserID     int64  `json:"user_id"`
	CampaignID int64  `json:"campaign_id"`
	Status     bool   `json:"status"`
}

func (d Domain) ToJson() []byte {
	js, _ := json.Marshal(d)
	return js
}

func (d Domain) ToJsonString() string {
	return string(d.ToJson())
}
