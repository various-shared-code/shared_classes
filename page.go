package classes

import "encoding/json"

type Page struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"desc"`
	Status      int8        `json:"status"`
	UserID      int64       `json:"user_id"`
	CampaignID  int64       `json:"campaign_id"`
	DomainID    int64       `json:"domain_id"`
	TemplateID  int64       `json:"template_id"`
	Content     PageContent `json:"page_content"`
}

func (p Page) ToJson() []byte {
	js, _ := json.Marshal(p)
	return js
}
