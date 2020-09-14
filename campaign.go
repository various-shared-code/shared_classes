package classes

import "encoding/json"

type Campaign struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	UserID      int64  `json:"user_id"`
	Status      int8   `json:"status"`
}

func (c Campaign) ToJson() []byte {
	js, _ := json.Marshal(c)
	return js
}
