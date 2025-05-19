package go_praxis

// ----------pre generate-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/cashier_api/cashier

type PraxisDepositReq struct {
	//MerchantID      string                     `json:"merchant_id"`
	//ApplicationKey  string                     `json:"application_key"`
	//Locale          string                     `json:"locale"`
	//Version         string                     `json:"version"`
	Intent          string                     `json:"intent"`
	Currency        string                     `json:"currency"`
	Amount          int                        `json:"amount"`
	Cid             string                     `json:"cid"`
	CustomerToken   string                     `json:"customer_token"`
	CustomerData    *PraxisDepositCustomerData `json:"customer_data"`
	PaymentMethod   string                     `json:"payment_method"`
	Gateway         string                     `json:"gateway"`
	ValidationURL   string                     `json:"validation_url"`
	NotificationURL string                     `json:"notification_url"`
	ReturnURL       string                     `json:"return_url"`
	OrderID         string                     `json:"order_id"`
	Timestamp       int64                      `json:"timestamp"`
}

type PraxisDepositCustomerData struct {
	Country   string `json:"country"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"` // Date of birth (format: YYYY-MM-DD)
	Email     string `json:"email"`
	Phone     string `json:"phone"`   // Should include country code
	Zip       string `json:"zip"`     // Postal/ZIP code
	State     string `json:"state"`   // State/Province
	City      string `json:"city"`    // City
	Address   string `json:"address"` // Street address
}

type PraxisDepositCardData struct {
	CardNumber string `json:"card_number"` // Encrypted card number (e.g., "ZMq4wDaiaQ/xOwMEcQ7R3ASjTnoOMu+avLuJYgAnz1Q=")
	CardExp    string `json:"card_exp"`    // Encrypted expiration date (e.g., "WI8V4bE5/l8fIhUv6aMO8w==")
	CVV        string `json:"cvv"`         // Encrypted CVV code
}

//--------------------------------------------------------

type PraxisDepositRsp struct {
	Status      int                        `json:"status"`
	Description string                     `json:"description"`
	RedirectURL string                     `json:"redirect_url"`
	Customer    *PraxisDevicePsRspCustomer `json:"customer"`
	Session     *PraxisDevicePsRspSession  `json:"session"`
	Version     string                     `json:"version"`
	Timestamp   int64                      `json:"timestamp"`
}

//---------------------------------------------

type PraxisWithdrawReq struct {
	//MerchantID      string                     `json:"merchant_id"`
	//ApplicationKey  string                     `json:"application_key"`
	//Locale          string                     `json:"locale"`
	//Version         string                     `json:"version"`
	Intent          string                     `json:"intent"`
	Currency        string                     `json:"currency"`
	Amount          int                        `json:"amount"`
	Balance         int                        `json:"balance"`
	Cid             string                     `json:"cid"`
	CustomerToken   string                     `json:"customer_token"`
	CustomerData    *PraxisDepositCustomerData `json:"customer_data"`
	CardData        *PraxisDepositCardData     `json:"card_data"`
	PaymentMethod   string                     `json:"payment_method"`
	Gateway         string                     `json:"gateway"`
	ValidationURL   string                     `json:"validation_url"`
	NotificationURL string                     `json:"notification_url"`
	ReturnURL       string                     `json:"return_url"`
	OrderID         string                     `json:"order_id"`
	Timestamp       int64                      `json:"timestamp"`
}

type PraxisWithdrawResp struct {
	Status      int                        `json:"status"`
	Description string                     `json:"description"`
	RedirectURL string                     `json:"redirect_url"`
	Customer    *PraxisDevicePsRspCustomer `json:"customer"`
	Session     *PraxisDevicePsRspSession  `json:"session"`
	Version     string                     `json:"version"`
	Timestamp   int64                      `json:"timestamp"`
}

// PraxisDevicePsRspCustomer represents customer data in response
type PraxisDevicePsRspCustomer struct {
	AVSAlert          int    `json:"avs_alert"`
	CustomerToken     string `json:"customer_token"`
	VerificationAlert int    `json:"verification_alert"`
}

// PraxisDevicePsRspSession represents session data in response
type PraxisDevicePsRspSession struct {
	Amount            float64 `json:"amount"`
	AuthToken         string  `json:"auth_token"`
	Cid               string  `json:"cid"`
	Currency          string  `json:"currency"`
	Intent            string  `json:"intent"`
	OrderID           string  `json:"order_id"`
	ProcessedAmount   float64 `json:"processed_amount"`
	ProcessedCurrency string  `json:"processed_currency"`
	SessionStatus     string  `json:"session_status"`
}

// ----------deposit callback-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/webhooks/notification

type PraxisBackReq struct {
	MerchantID     string                        `json:"merchant_id"`
	ApplicationKey string                        `json:"application_key"`
	Customer       *PraxisBackReqCustomerData    `json:"customer"`
	Session        *PraxisBackReqSessionData     `json:"session"`
	Transaction    *PraxisBackReqTransactionData `json:"transaction"`
	Version        string                        `json:"version"`
	Timestamp      int64                         `json:"timestamp"`
}

type PraxisBackReqCustomerData struct {
	CustomerToken     string `json:"customer_token"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	AVSAlert          *int   `json:"avs_alert,omitempty"`          // Using pointer to allow null/nil
	VerificationAlert string `json:"verification_alert,omitempty"` // omitempty if null
}

type PraxisBackReqSessionData struct {
	AuthToken         string  `json:"auth_token"`
	Intent            string  `json:"intent"`
	SessionStatus     string  `json:"session_status"`
	OrderID           string  `json:"order_id"`
	Currency          string  `json:"currency"`
	Amount            float64 `json:"amount"`
	ConversionRate    float64 `json:"conversion_rate"`
	ProcessedCurrency string  `json:"processed_currency"`
	ProcessedAmount   float64 `json:"processed_amount"`
	PaymentMethod     string  `json:"payment_method"`
	Gateway           string  `json:"gateway"`
	Pin               string  `json:"pin,omitempty"`       // omitempty if empty
	Variable1         string  `json:"variable1,omitempty"` // omitempty if empty
	Variable2         string  `json:"variable2,omitempty"` // omitempty if empty
	Variable3         string  `json:"variable3,omitempty"` // omitempty if empty
}

type PraxisBackReqTransactionData struct {
	TransactionType   string                 `json:"transaction_type"`
	TransactionStatus string                 `json:"transaction_status"`
	Tid               string                 `json:"tid"`
	TransactionID     string                 `json:"transaction_id"`
	Currency          string                 `json:"currency"`
	Amount            string                 `json:"amount"`
	ConversionRate    string                 `json:"conversion_rate"`
	ProcessedCurrency string                 `json:"processed_currency"`
	ProcessedAmount   string                 `json:"processed_amount"`
	Fee               string                 `json:"fee"`
	FeeIncluded       string                 `json:"fee_included"`
	FeeType           string                 `json:"fee_type"`
	PaymentMethod     string                 `json:"payment_method"`
	PaymentProcessor  string                 `json:"payment_processor"`
	Gateway           string                 `json:"gateway"`
	Card              *PraxisBackReqCardData `json:"card,omitempty"`
}

// PraxisDepositBackReqCardData represents card data in transaction
type PraxisBackReqCardData struct {
	CardToken      string `json:"card_token"`
	CardType       string `json:"card_type"`
	CardNumber     string `json:"card_number"`
	CardExp        string `json:"card_exp"`
	CardIssuerName string `json:"card_issuer_name"`
	CardHolder     string `json:"card_holder"`
}

// =============
// 返回给三方的
type PraxisBackResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Timestamp   int64  `json:"timestamp"`
}
