package cloudpayments

type CPResponse[T any] struct {
	Model   T      `json:"Model"`
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type Transaction struct {
	ReasonCode                            int                    `json:"ReasonCode"`
	PublicId                              string                 `json:"PublicId"`
	TerminalUrl                           string                 `json:"TerminalUrl"`
	TransactionId                         int                    `json:"TransactionId"`
	Amount                                int                    `json:"Amount"`
	Currency                              string                 `json:"Currency"`
	CurrencyCode                          int                    `json:"CurrencyCode"`
	PaymentAmount                         int                    `json:"PaymentAmount"`
	PaymentCurrency                       string                 `json:"PaymentCurrency"`
	PaymentCurrencyCode                   int                    `json:"PaymentCurrencyCode"`
	InvoiceId                             string                 `json:"InvoiceId"`
	AccountId                             string                 `json:"AccountId"`
	Email                                 string                 `json:"Email"`
	Description                           string                 `json:"Description"`
	JsonData                              map[string]interface{} `json:"JsonData"`
	CreatedDate                           string                 `json:"CreatedDate"`
	PayoutDate                            string                 `json:"PayoutDate"`
	PayoutDateIso                         string                 `json:"PayoutDateIso"`
	PayoutAmount                          interface{}            `json:"PayoutAmount"`
	CreatedDateIso                        string                 `json:"CreatedDateIso"`
	AuthDate                              string                 `json:"AuthDate"`
	AuthDateIso                           string                 `json:"AuthDateIso"`
	ConfirmDate                           interface{}            `json:"ConfirmDate"`
	ConfirmDateIso                        interface{}            `json:"ConfirmDateIso"`
	AuthCode                              string                 `json:"AuthCode"`
	TestMode                              bool                   `json:"TestMode"`
	Rrn                                   interface{}            `json:"Rrn"`
	OriginalTransactionId                 interface{}            `json:"OriginalTransactionId"`
	FallBackScenarioDeclinedTransactionId interface{}            `json:"FallBackScenarioDeclinedTransactionId"`
	IpAddress                             string                 `json:"IpAddress"`
	IpCountry                             string                 `json:"IpCountry"`
	IpCity                                string                 `json:"IpCity"`
	IpRegion                              string                 `json:"IpRegion"`
	IpDistrict                            string                 `json:"IpDistrict"`
	IpLatitude                            float64                `json:"IpLatitude"`
	IpLongitude                           float64                `json:"IpLongitude"`
	CardFirstSix                          string                 `json:"CardFirstSix"`
	CardLastFour                          string                 `json:"CardLastFour"`
	CardExpDate                           string                 `json:"CardExpDate"`
	CardType                              string                 `json:"CardType"`
	CardProduct                           string                 `json:"CardProduct"`
	CardCategory                          string                 `json:"CardCategory"`
	EscrowAccumulationId                  interface{}            `json:"EscrowAccumulationId"`
	IssuerBankCountry                     string                 `json:"IssuerBankCountry"`
	Issuer                                string                 `json:"Issuer"`
	CardTypeCode                          int                    `json:"CardTypeCode"`
	Status                                string                 `json:"Status"`
	StatusCode                            int                    `json:"StatusCode"`
	CultureName                           string                 `json:"CultureName"`
	Reason                                string                 `json:"Reason"`
	CardHolderMessage                     string                 `json:"CardHolderMessage"`
	Type                                  int                    `json:"Type"`
	Refunded                              bool                   `json:"Refunded"`
	Name                                  string                 `json:"Name"`
	Token                                 string                 `json:"Token"`
	SubscriptionId                        interface{}            `json:"SubscriptionId"`
	GatewayName                           string                 `json:"GatewayName"`
	ApplePay                              bool                   `json:"ApplePay"`
	AndroidPay                            bool                   `json:"AndroidPay"`
	WalletType                            string                 `json:"WalletType"`
	TotalFee                              int                    `json:"TotalFee"`
}

type Secure3D struct {
	TransactionId        int         `json:"TransactionId"`
	PaReq                string      `json:"PaReq"`
	GoReq                interface{} `json:"GoReq"`
	AcsUrl               string      `json:"AcsUrl"`
	ThreeDsSessionData   interface{} `json:"ThreeDsSessionData"`
	IFrameIsAllowed      bool        `json:"IFrameIsAllowed"`
	FrameWidth           interface{} `json:"FrameWidth"`
	FrameHeight          interface{} `json:"FrameHeight"`
	ThreeDsCallbackId    string      `json:"ThreeDsCallbackId"`
	EscrowAccumulationId interface{} `json:"EscrowAccumulationId"`
}

type Payer struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	MiddleName string `json:"MiddleName"`
	Birth      string `json:"Birth"`
	Street     string `json:"Street"`
	Address    string `json:"Address"`
	City       string `json:"City"`
	Country    string `json:"Country"`
	Phone      string `json:"Phone"`
	Postcode   string `json:"Postcode"`
}

type ChargeCardInput struct {
	CardCryptogramPacket string
	Amount               int
	IpAddress            string
	Currency             Currency
	Name                 string
	PaymentUrl           string
	InvoiceId            string
	Description          string
	CultureName          string
	AccountId            string
	Email                string
	Payer                Payer
	JsonData             map[string]interface{}
	RequireConfirmation  bool
}
