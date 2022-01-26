package entities

type User struct {
	ID          string `json:"id,omitempty"`
	AccountCode string `json:"account_code,omitempty"`
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
}

// RequestGetAllUser defines the field which should be the request payload.
type RequestGetAllUser struct {
}

type ResponseGetAllUser struct {
	Status string `json:"status,omitempty"`
	Data   User
}
