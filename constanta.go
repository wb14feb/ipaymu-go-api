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
