package go_neteller

// ----------pre generate-------------------------

// 充值/提现是同一个request, 同一个response.  只是里边一个参数不同而已.
type NetellerPaymentReq struct {
	OutType        int     `json:"outType"` //通过这个参数区别是deposit还是withdraw
	MerchantRefNum string  `json:"merchantRefNum"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	UserId         int64   `json:"userId"`
	Email          string  `json:"email"` //要去user里拿
}

type NetellerPaymentRsp struct {
	ID                 string    `json:"id"`
	PaymentHandleToken string    `json:"paymentHandleToken"`
	MerchantRefNum     string    `json:"merchantRefNum"`
	Links              []LinksBO `json:"links"`
	Status             string    `json:"status"`
}

type LinksBO struct {
	Rel  string `json:"rel"`  // Relation type of the link
	Href string `json:"href"` // URL of the link
}

// ---------- callback-------------------------

type NetellerBackReq struct {
	EventType     string    `json:"eventType"`
	AttemptNumber string    `json:"attemptNumber"`
	ResourceId    string    `json:"resourceId"`
	EventDate     string    `json:"eventDate"`
	Mode          string    `json:"mode"`
	EventName     string    `json:"eventName"`
	Links         []LinksBO `json:"links"`
	Payload       PayLoadBO `json:"payload"`
}

type PayLoadBO struct {
	AccountId      string  `json:"accountId"`
	Id             string  `json:"id"`
	MerchantRefNum string  `json:"merchantRefNum"`
	Amount         float64 `json:"amount"` // Using big.Float for precise decimal representation
	CurrencyCode   string  `json:"currencyCode"`
	Status         string  `json:"status"`
	PaymentType    string  `json:"paymentType"`
	TxnTime        string  `json:"txnTime"`
}

//----------------------------

type NetellerProcessReq struct {
	MerchantRefNum     string  `json:"merchantRefNum"`
	Amount             float64 `json:"amount"`
	CurrencyCode       string  `json:"currencyCode"`
	PaymentHandleToken string  `json:"paymentHandleToken"`
	OutType            int     `json:"outType"`
}

type NetellerProcessRsp struct {
	ID     string `json:"id"`     // Transaction ID
	Status string `json:"status"` // Payment status
}
