package ipaymu_go_api

type RequestDirectMaster struct {
	Name          string        `json:"name"`
	Phone         string        `json:"phone"`
	Email         string        `json:"email"`
	Amount        float64       `json:"amount"`
	NotifyUrl     string        `json:"notifyUrl"`
	Expired       string        `json:"expired"`
	ExpiredType   ExpiredType   `json:"expiredType"`
	Comments      string        `json:"comments"`
	ReferenceId   string        `json:"referenceId"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
}

func (r *RequestDirectMaster) AddBuyer(name, phone, email string) {
	r.Name = name
	r.Phone = phone
	r.Email = email
}

type Product struct {
	Product []string  `json:"product,omitempty"`
	Qty     []int8    `json:"qty,omitempty"`
	Price   []float64 `json:"price,omitempty"`
}

type ProductWithCOD struct {
	Product
	Weight []float32 `json:"weight,omitempty"`
	Width  []float32 `json:"width,omitempty"`
	Height []float32 `json:"height,omitempty"`
	Length []float32 `json:"length,omitempty"`
}

type RequestDirectVA struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelVA `json:"paymentChannel"`
	Product
}

func NewRequestDirectVA(channel PaymentChannelVA) *RequestDirectVA {
	req := &RequestDirectVA{}
	req.PaymentMethod = VirtualAccount
	req.PaymentChannel = channel
	return req
}

type RequestDirectCOD struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelCOD `json:"paymentChannel"`
	*ProductWithCOD
	DeliveryArea    string `json:"deliveryArea"`
	DeliveryAddress string `json:"deliveryAddress"`
}

func NewRequestDirectCOD(codChannel PaymentChannelCOD) *RequestDirectCOD {
	req := &RequestDirectCOD{}
	req.PaymentMethod = COD
	req.PaymentChannel = codChannel
	return req
}

type RequestDirectConStore struct {
	RequestDirectMaster
	PaymentChannel PaymentChannelConStore `json:"paymentChannel"`
}

type RequestRedirect struct {
	Description   string         `json:"description"`
	ReturnUrl     string         `json:"returnUrl"`
	NotifyUrl     string         `json:"notifyUrl"`
	CancelUrl     string         `json:"cancelUrl"`
	ReferenceId   *string        `json:"referenceId"`
	Product       []string       `json:"product"`
	Qty           []int8         `json:"qty"`
	Price         []float64      `json:"price"`
	Weight        []float32      `json:"weight"`
	Dimension     []string       `json:"dimension"`
	BuyerName     *string        `json:"buyerName"`
	BuyerEmail    *string        `json:"buyerEmail"`
	BuyerPhone    *string        `json:"buyerPhone"`
	PickupArea    *string        `json:"pickupArea"`
	PickupAddress *string        `json:"pickupAddress"`
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
}

func NewRequestRedirect() *RequestRedirect {
	return &RequestRedirect{}
}

func (r *RequestRedirect) AddProduct(product string, qty int8, price float64, weight *float32, dimension *string) {
	r.Product = append(r.Product, product)
	r.Qty = append(r.Qty, qty)
	r.Price = append(r.Price, price)
	if weight != nil {
		r.Weight = append(r.Weight, *weight)
	}
	if dimension != nil {
		r.Dimension = append(r.Dimension, *dimension)
	}
}
