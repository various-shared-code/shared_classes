package classes

import "encoding/json"

type PageContent struct {
	Language  string `json:"language"`
	SessionID string `json:"session_id"`

	Title      string              `json:"title"`
	Logo       string              `json:"logo"`
	Components []EditableComponent `json:"components"`
}

func (p PageContent) ToJson() []byte {
	js, _ := json.Marshal(p)
	return js
}

func (p PageContent) ToJsonString() string {
	return string(p.ToJson())
}
