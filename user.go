package classes

import "encoding/json"

type AccountType string

type User struct {
	ID                   int64       `json:"id"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	Email                string      `json:"email"`
	Type                 AccountType `json:"type"`
	RegistrationDateTime int64       `json:"registration_date_time"`
	LoginDateTime        int64       `json:"login_date_time"`
	Status               int8        `json:"status"`
	AllowsMarketing      bool        `json:"allows_marketing"`
}

func (u User) ToJson() []byte {
	js, _ := json.Marshal(u)
	return js
}

func (u User) ToJsonString() string {
	return string(u.ToJson())
}
