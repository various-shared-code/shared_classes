package classes

import "encoding/json"

type PageContent struct {
	Language  string `json:"language"`
	SessionID string `json:"session_id"`

	Title           string `json:"title"`
	Logo            string `json:"logo"`
	MainSlogan      string `json:"main_slogan"`
	MainDescription string `json:"main_description"`
}

func (p PageContent) ToJson() []byte {
	js, _ := json.Marshal(p)
	return js
}
