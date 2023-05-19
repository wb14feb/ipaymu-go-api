package ipaymu_go_api

type PaymentMethod string

const (
	VirtualAccount   PaymentMethod = "va"
	ConvenienceStore PaymentMethod = "cstore"
	COD              PaymentMethod = "cod"
	QRISMethod       PaymentMethod = "qris"
	Paylater         PaymentMethod = "paylater"
)

type PaymentChannelVA string

const (
	BAG       PaymentChannelVA = "bag"
	BCA       PaymentChannelVA = "bca"
	BNI       PaymentChannelVA = "bni"
	CimbNiaga PaymentChannelVA = "cimb"
	Mandiri   PaymentChannelVA = "mandiri"
	Muamalat  PaymentChannelVA = "bmi"
	BRI       PaymentChannelVA = "bri"
	BSI       PaymentChannelVA = "bsi"
	Permata   PaymentChannelVA = "permata"
	Danamon   PaymentChannelVA = "danamon"
)

type PaymentChannelConStore string

const (
	Alfamart  PaymentChannelConStore = "alfamart"
	Indomaret PaymentChannelConStore = "indomaret"
)

type PaymentChannelCOD string

const RPX PaymentChannelCOD = "rpx"

type PaymentChannelPaylater string

const Akulaku PaymentChannelPaylater = "akulaku"

type PaymentChannelQRIS string

const QRIS PaymentChannelQRIS = "qris"

type ExpiredType string

const (
	Days    ExpiredType = "days"
	Hours   ExpiredType = "hours"
	Minutes ExpiredType = "minutes"
	Seconds ExpiredType = "seconds"
)

type EnvironmentType string

const (
	Production EnvironmentType = "https://my.ipaymu.com"
	Sandbox    EnvironmentType = "https://sandbox.ipaymu.com"
)

type PaymentStatus int8

const (
	Expired         PaymentStatus = -2
	Pending         PaymentStatus = 0
	Success         PaymentStatus = 1
	Cancel          PaymentStatus = 2
	Refund          PaymentStatus = 3
	Error           PaymentStatus = 4
	Failed          PaymentStatus = 5
	SuccessUnsettle PaymentStatus = 6
	Escrow          PaymentStatus = 7
)

type FilterDate string

const (
	CreatedAt FilterDate = "created_at"
	PaidAt    FilterDate = "indate"
)

type PaymentType int8

const (
	Mutasi          PaymentType = 0
	Withdraw        PaymentType = 2
	Commission      PaymentType = 5
	BankTransferVA  PaymentType = 7
	CashOnDelivery  PaymentType = 10
	ConStore        PaymentType = 11
	MoveBalance     PaymentType = 20
	AkulakuPaylater PaymentType = 22
	QRISReceiver    PaymentType = 23
	EDC             PaymentType = 24
	QRISPayment     PaymentType = 27
)

type FilterOrderBy string

const (
	ID     FilterOrderBy = "id"
	Create FilterOrderBy = "created_at"
	Paid   FilterOrderBy = "indate"
)

type FilterOrder string

const (
	ASC  FilterOrder = "ASC"
	DESC FilterOrder = "DESC"
)

type FilterLanguage string

const (
	EN  FilterLanguage = "en"
	IDN FilterLanguage = "id"
)

type FilterLockStatus int8

const (
	Locked   FilterLockStatus = 1
	Unlocked FilterLockStatus = 0
)
