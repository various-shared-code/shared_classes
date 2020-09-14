package classes

import "encoding/json"

type TemplateDescriptor struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Preview     string   `json:"preview"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`

	Content TemplateContent `json:"content"`
}

func (d TemplateDescriptor) ToJson() []byte {
	js, _ := json.Marshal(d)
	return js
}

func (d TemplateDescriptor) ToJsonString() string {
	return string(d.ToJson())
}
