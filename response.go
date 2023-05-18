package ipaymu_go_api

type Response struct {
	Status int64
	
	// v2 payment
	Message string        `json:"Message,omitempty"`
	Data    *ResponseData `json:"Data,omitempty"`
}

type ResponseData struct {
	SessionId     string `json:"SessionId,omitempty"`
	TransactionId int64  `json:"TransactionId,omitempty"`
	ReferenceId   string `json:"ReferenceId,omitempty"`
	Via           string `json:"Via,omitempty"`
	Channel       string `json:"Channel,omitempty"`
	PaymentNo     string `json:"PaymentNo,omitempty"`
	PaymentName   string `json:"PaymentName,omitempty"`
	Total         string `json:"Total,omitempty"`
	Fee           int64  `json:"Fee,omitempty"`
	Expired       string `json:"Expired,omitempty"`
	Note          string `json:"Note,omitempty"`
	Url           string `json:"Url,omitempty"`
}
