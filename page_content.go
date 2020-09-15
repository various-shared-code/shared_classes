package classes

import "encoding/json"

type PageContent struct {
	Language  string `json:"language"`
	SessionID string `json:"session_id"`

	Title string `json:"title"`
	Logo  string `json:"logo"`

	MainHighlight   string `json:"main_highlight"`
	MainDescription string `json:"main_description"`
}

func (p PageContent) ToJson() []byte {
	js, _ := json.Marshal(p)
	return js
}

func (p PageContent) ToJsonString() string {
	return string(p.ToJson())
}
