package classes

import "encoding/json"

type EditableComponent struct {
	Type       string `json:"type"`
	HtmlId     string `json:"htmlId"`
	FontFace   string `json:"fontFace"`
	FontSize   int    `json:"fontSize"`
	FontWeight string `json:"fontWeight"`
	Value      string `json:"value"`
}

func (c EditableComponent) ToJson() []byte {
	js, _ := json.Marshal(c)
	return js
}

func (c EditableComponent) ToJsonString() string {
	return string(c.ToJson())
}
