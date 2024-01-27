package ipaymu_go_api

type Response struct {
	Status int64 `json:"status"`

	// v2 payment
	Message string        `json:"message,omitempty"`
	Data    *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	SessionId     string  `json:"sessionId,omitempty"`
	TransactionId int64   `json:"transactionId,omitempty"`
	ReferenceId   string  `json:"referenceId,omitempty"`
	Via           string  `json:"via,omitempty"`
	Channel       string  `json:"channel,omitempty"`
	PaymentNo     string  `json:"paymentNo,omitempty"`
	PaymentName   string  `json:"paymentName,omitempty"`
	Total         float64 `json:"total,omitempty"`
	Fee           int64   `json:"fee,omitempty"`
	Expired       string  `json:"expired,omitempty"`
	Note          string  `json:"note,omitempty"`
	Url           string  `json:"url,omitempty"`
}

// intermidiate struct for unmarshalling, you can ignore this
type ResponseUpper struct {
	Status int64 `json:"Status"`

	// v2 payment
	Message string             `json:"Message,omitempty"`
	Data    *ResponseDataUpper `json:"Data,omitempty"`
}

// intermidiate struct for unmarshalling, you can ignore this
type ResponseDataUpper struct {
	SessionId     string  `json:"SessionId,omitempty"`
	TransactionId int64   `json:"TransactionId,omitempty"`
	ReferenceId   string  `json:"ReferenceId,omitempty"`
	Via           string  `json:"Via,omitempty"`
	Channel       string  `json:"Channel,omitempty"`
	PaymentNo     string  `json:"PaymentNo,omitempty"`
	PaymentName   string  `json:"PaymentName,omitempty"`
	Total         float64 `json:"Total,omitempty"`
	Fee           int64   `json:"Fee,omitempty"`
	Expired       string  `json:"Expired,omitempty"`
	Note          string  `json:"Note,omitempty"`
	Url           string  `json:"Url,omitempty"`
}

func (ru *ResponseUpper) ToResponse() Response {
	return Response{
		Status:  ru.Status,
		Message: ru.Message,
		Data:    ru.Data.ToResponseData(),
	}
}

func (rdu *ResponseDataUpper) ToResponseData() *ResponseData {
	if rdu == nil {
		return nil
	}
	return &ResponseData{
		SessionId:     rdu.SessionId,
		TransactionId: rdu.TransactionId,
		ReferenceId:   rdu.ReferenceId,
		Via:           rdu.Via,
		Channel:       rdu.Channel,
		PaymentNo:     rdu.PaymentNo,
		PaymentName:   rdu.PaymentName,
		Total:         rdu.Total,
		Fee:           rdu.Fee,
		Expired:       rdu.Expired,
		Note:          rdu.Note,
		Url:           rdu.Url,
	}
}

type ResponseCheck struct {
	Status int `json:"Status"`
	Data   struct {
		TransactionId  int         `json:"transactionId"`
		SessionId      string      `json:"sessionId"`
		ReferenceId    interface{} `json:"referenceId"`
		RelatedId      int         `json:"relatedId"`
		Sender         string      `json:"sender"`
		Receiver       string      `json:"receiver"`
		Amount         int         `json:"amount"`
		Fee            int         `json:"fee"`
		Status         int         `json:"status"`
		StatusDesc     string      `json:"statusDesc"`
		Type           int         `json:"type"`
		TypeDesc       string      `json:"typeDesc"`
		Notes          string      `json:"notes"`
		CreatedDate    string      `json:"createdDate"`
		ExpiredDate    string      `json:"expiredDate"`
		SuccessDate    string      `json:"successDate"`
		SettlementDate string      `json:"settlementDate"`
	} `json:"data"`
	Message string `json:"message"`
}

type ResponseBalance struct {
	Status int `json:"status"`
	Data   struct {
		Va              string  `json:"va"`
		MerchantBalance float64 `json:"merchantBalance"`
		MemberBalance   float64 `json:"memberBalance"`
	} `json:"data"`
	Message string `json:"message"`
}

type ResponseTransaction struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Transaction []Transaction `json:"transaction"`
		Pagination  struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"Pagination"`
	} `json:"Data"`
}

type Transaction struct {
	TransactionId  int         `json:"TransactionId"`
	SessionId      *string     `json:"SessionId"`
	ReferenceId    *string     `json:"ReferenceId"`
	RelatedId      int         `json:"RelatedId"`
	Sender         string      `json:"Sender"`
	Receiver       string      `json:"Receiver"`
	Amount         int         `json:"Amount"`
	Fee            int         `json:"Fee"`
	Status         int         `json:"Status"`
	StatusDesc     string      `json:"StatusDesc"`
	PaidStatus     string      `json:"PaidStatus"`
	Type           int         `json:"Type"`
	TypeDesc       string      `json:"TypeDesc"`
	Notes          *string     `json:"Notes"`
	IsEscrow       bool        `json:"IsEscrow"`
	CreatedDate    string      `json:"CreatedDate"`
	ExpiredDate    string      `json:"ExpiredDate"`
	SuccessDate    interface{} `json:"SuccessDate"`
	SettlementDate interface{} `json:"SettlementDate"`
	PaymentChannel string      `json:"PaymentChannel"`
	PaymentCode    string      `json:"PaymentCode"`
	BuyerName      string      `json:"BuyerName"`
	BuyerPhone     string      `json:"BuyerPhone"`
	BuyerEmail     string      `json:"BuyerEmail"`
}

type ResponseListPayment struct {
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
	Data    []struct {
		Code        string                 `json:"Code"`
		Name        string                 `json:"Name"`
		Description string                 `json:"Description"`
		Channels    []PaymentChannelDetail `json:"Channels,omitempty"`
	} `json:"Data"`
}

type ResponseListPaymentLower struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []struct {
		Code        string                      `json:"code"`
		Name        string                      `json:"name"`
		Description string                      `json:"description"`
		Channels    []PaymentChannelDetailLower `json:"channels,omitempty"`
	} `json:"data"`
}

type PaymentChannelDetail struct {
	Code                 string `json:"Code"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	PaymentIntrucionsDoc string `json:"PaymentIntrucionsDoc"`
	TransactionFee       struct {
		ActualFee     float64 `json:"ActualFee" example:"0.0"  description:"actual fee, can be percent"`
		ActualFeeType string  `json:"ActualFeeType"`
		AdditionalFee int     `json:"AdditionalFee"`
	}
}

type PaymentChannelDetailLower struct {
	Code                   string `json:"code"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	PaymentInstructionsDoc string `json:"paymentInstructionsDoc"`
	TransactionFee         struct {
		ActualFee     float64 `json:"actualFee"  example:"0.0"  description:"actual fee, can be percent"`
		ActualFeeType string  `json:"actualFeeType" examples:"PERCENT, FLAT"  description:"actual fee type, can be PERCENT or FLAT"`
		AdditionalFee int     `json:"additionalFee"`
	}
}

func (r ResponseListPayment) EncodeJsonLowerCase() ResponseListPaymentLower {
	var res ResponseListPaymentLower
	res.Status = r.Status
	res.Success = r.Success
	res.Message = r.Message

	for _, v := range r.Data {
		var data struct {
			Code        string                      `json:"code"`
			Name        string                      `json:"name"`
			Description string                      `json:"description"`
			Channels    []PaymentChannelDetailLower `json:"channels,omitempty"`
		}
		data.Code = v.Code
		data.Name = v.Name
		data.Description = v.Description
		var channels []PaymentChannelDetailLower

		for _, strct := range v.Channels {
			channels = append(channels, strct.EncodeJsonLowerCase())
		}

		data.Channels = channels

		res.Data = append(res.Data, data)
	}
	return res
}

func (r PaymentChannelDetail) EncodeJsonLowerCase() PaymentChannelDetailLower {
	var res PaymentChannelDetailLower
	res.Code = r.Code
	res.Name = r.Name
	res.Description = r.Description
	res.PaymentInstructionsDoc = r.PaymentIntrucionsDoc
	res.TransactionFee.ActualFee = r.TransactionFee.ActualFee
	res.TransactionFee.ActualFeeType = r.TransactionFee.ActualFeeType
	res.TransactionFee.AdditionalFee = r.TransactionFee.AdditionalFee

	return res

}
