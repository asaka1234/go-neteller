package go_neteller

type NetellerInitParams struct {
	MerchantId  string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`     // merchantId
	MerchantKey string `json:"merchantKey" mapstructure:"merchantKey" config:"merchantKey"  yaml:"merchantKey"` // accessKey

	CreatePaymentHandleUrl      string `json:"createPaymentHandleUrl" mapstructure:"createPaymentHandleUrl" config:"createPaymentHandleUrl"  yaml:"createPaymentHandleUrl"`
	ProcessStandaloneCreditsUrl string `json:"processStandaloneCreditsUrl" mapstructure:"processStandaloneCreditsUrl" config:"processStandaloneCreditsUrl"  yaml:"processStandaloneCreditsUrl"`
	ProcessPaymentsUrl          string `json:"processPaymentsUrl" mapstructure:"processPaymentsUrl" config:"processPaymentsUrl"  yaml:"processPaymentsUrl"`
	GetPaymentHandleUrl         string `json:"getPaymentHandleUrl" mapstructure:"getPaymentHandleUrl" config:"getPaymentHandleUrl"  yaml:"getPaymentHandleUrl"`

	PaymentFeBackUrl string `json:"paymentFeBackUrl" mapstructure:"paymentFeBackUrl" config:"paymentFeBackUrl"  yaml:"paymentFeBackUrl"` //前端回跳地址
}

// ----------pre generate-------------------------

type NetellerPaymentHandleReq struct {
	MerchantRefNum string         `json:"merchantRefNum" mapstructure:"merchantRefNum"` //这里需要放merchantOrderNo
	Amount         int            `json:"amount" mapstructure:"amount"`                 //这里需要做单位转换,用的是法币的最小单位. 比如1.2美元，这里传的是120美分
	CurrencyCode   string         `json:"currencyCode" mapstructure:"currencyCode"`     //币种
	Neteller       NetellerDetail `json:"neteller" mapstructure:"neteller"`
	//sdk搞定
	//TransactionType string         `json:"transactionType" mapstructure:"transactionType"` //枚举: PAYMENT (付款给商户)， STANDALONE_CREDIT (商户付给User)
	//PaymentType string `json:"paymentType" mapstructure:"paymentType"` //Fixed  NETELLER
	//ReturnLinks []ReturnLink `json:"returnLinks" mapstructure:"returnLinks"` //设置回调地址
}

type NetellerDetail struct {
	ConsumerId string `json:"consumerId" mapstructure:"consumerId"` //收款人的邮箱 This is the email address of the customer who is receiving the payment
	//option
	//ConsumerIdLocked   bool   `json:"consumerIdLocked" mapstructure:"consumerIdLocked"`
	//Detail1Description string `json:"detail1Description" mapstructure:"detail1Description"`
	//Detail1Text        string `json:"detail1Text" mapstructure:"detail1Text"`
}

type BillingDetails struct {
	Street  string `json:"street" mapstructure:"street"` //账单地址
	Street2 string `json:"street2" mapstructure:"street2"`
	City    string `json:"city" mapstructure:"city"`
	Zip     string `json:"zip" mapstructure:"zip"`
	Country string `json:"country" mapstructure:"country"`
}

// 请求时指定回调地址
type ReturnLink struct {
	Rel  string `json:"rel" mapstructure:"rel"`   //This is the link type, 枚举: default,on_completed,on_failed,on_cancelled
	Href string `json:"href" mapstructure:"href"` //The actual URL
}

//-------------------------------------

type NetellerProcessStandaloneCreditsReq struct {
	MerchantRefNum     string `json:"merchantRefNum" mapstructure:"merchantRefNum"`
	Amount             int    `json:"amount" mapstructure:"amount"`
	CurrencyCode       string `json:"currencyCode" mapstructure:"currencyCode"`
	PaymentHandleToken string `json:"paymentHandleToken" mapstructure:"paymentHandleToken"`
}

type NetellerProcessStandaloneCreditsResp struct {
	Id                 string `json:"id"`          //psp的订单号
	PaymentType        string `json:"paymentType"` //写死的 NETELLER
	PaymentHandleToken string `json:"paymentHandleToken"`
	MerchantRefNum     string `json:"merchantRefNum"` //商户订单号
	CurrencyCode       string `json:"currencyCode"`
	Status             string `json:"status"` //枚举: RECEIVED,COMPLETED,HELD,FAILED,CANCELLED,PENDING
	Amount             int    `json:"amount"` //最小单位，比如美分
	ReturnLinks        []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"returnLinks"`
	LiveMode        bool `json:"liveMode"`
	GatewayResponse struct {
		Id              string `json:"id"`
		Amount          string `json:"amount"`
		Currency        string `json:"currency"`
		TransactionType string `json:"transactionType"`
		Description     string `json:"description"`
		Status          string `json:"status"`
		Processor       string `json:"processor"`
	} `json:"gatewayResponse"`
	Neteller  NetellerDetail `json:"neteller"`
	BodyError ErrorInfo      `json:"bodyError"`
}

//-------------------------------------

type NetellerProcessPaymentsReq struct {
	MerchantRefNum     string `json:"merchantRefNum" mapstructure:"merchantRefNum"`
	Amount             int    `json:"amount" mapstructure:"amount"`
	CurrencyCode       string `json:"currencyCode" mapstructure:"currencyCode"`
	PaymentHandleToken string `json:"paymentHandleToken" mapstructure:"paymentHandleToken"`
}

type NetellerProcessPaymentsResp struct {
	Id                 string `json:"id"`
	PaymentType        string `json:"paymentType"`
	PaymentHandleToken string `json:"paymentHandleToken"`
	MerchantRefNum     string `json:"merchantRefNum"`
	CurrencyCode       string `json:"currencyCode"`
	SettleWithAuth     bool   `json:"settleWithAuth"`
	TxnTime            string `json:"txnTime"`
	BillingDetails     struct {
		Street1 string `json:"street1"`
		Street2 string `json:"street2"`
		City    string `json:"city"`
		Zip     string `json:"zip"`
		Country string `json:"country"`
	} `json:"billingDetails"`
	Status                  string `json:"status"`
	GatewayReconciliationId string `json:"gatewayReconciliationId"`
	Amount                  int    `json:"amount"`
	ConsumerIp              string `json:"consumerIp"`
	LiveMode                bool   `json:"liveMode"`
	GatewayResponse         struct {
		OrderId           string      `json:"orderId"`
		MerchantRefId     string      `json:"merchantRefId"`
		TotalAmount       interface{} `json:"totalAmount"`
		Currency          string      `json:"currency"`
		Lang              string      `json:"lang"`
		CustomerId        string      `json:"customerId"`
		VerificationLevel string      `json:"verificationLevel"`
		TransactionId     string      `json:"transactionId"`
		TransactionType   string      `json:"transactionType"`
		Description       string      `json:"description"`
		Status            string      `json:"status"`
		Processor         string      `json:"processor"`
	} `json:"gatewayResponse"`
	AvailableToSettle int `json:"availableToSettle"`
	Neteller          struct {
		ConsumerId       string `json:"consumerId"`
		ConsumerIdLocked bool   `json:"consumerIdLocked"`
	} `json:"neteller"`
	Settlements struct {
		Amount         int    `json:"amount"`
		TxnTime        string `json:"txnTime"`
		MerchantRefNum string `json:"merchantRefNum"`
		Id             string `json:"id"`
		Status         string `json:"status"`
	} `json:"settlements"`
}

//-----------------------------------------------------------

type NetellerPaymentHandleResp struct {
	//error
	Error ErrorInfo `json:"error"` //psp三方的订单号
	//succeed
	ID                 string `json:"id" mapstructure:"id"`                                 //psp的订单号
	PaymentHandleToken string `json:"paymentHandleToken" mapstructure:"paymentHandleToken"` //类似session
	MerchantRefNum     string `json:"merchantRefNum" mapstructure:"merchantRefNum"`
	CurrencyCode       string `json:"currencyCode" mapstructure:"currencyCode"`
	Status             string `json:"status" mapstructure:"status"` //枚举: PAYABLE,INITIATED,FAILED,EXPIRED,COMPLETED

	Amount          int             `json:"amount" mapstructure:"amount"` //转为最小法币单位后的
	GatewayResponse GatewayResponse `json:"gatewayResponse" mapstructure:"gatewayResponse"`
	Neteller        NetellerDetail  `json:"neteller" mapstructure:"neteller"`
	ReturnLinks     []ReturnLink    `json:"returnLinks" mapstructure:"returnLinks"`
	Links           []Link          `json:"links" mapstructure:"links"`
}

// 重要
type GatewayResponse struct {
	OrderId     string      `json:"orderId" mapstructure:"orderId"`         //psp的订单号
	TotalAmount interface{} `json:"totalAmount" mapstructure:"totalAmount"` //The total amount due for this order, including all items, fees, taxes
	Currency    string      `json:"currency" mapstructure:"currency"`
	Status      string      `json:"status" mapstructure:"status"` //枚举: pending,cancelled,failed (The order was not paid).,paid,expired (The order had expired. Default expiry time is 15 mins).
	Lang        string      `json:"lang" mapstructure:"lang"`     //en_US
	//Processor   string `json:"processor" mapstructure:"processor"` //固定: NETELLER
}

type Link struct {
	Rel  string `json:"rel" mapstructure:"rel"` //实际付款地址， 枚举: redirect_payment,default,first,prev,next,last。
	Href string `json:"href" mapstructure:"href"`
}

// -------------error--
type ErrorInfo struct {
	Code    string   `json:"code"`    //5279
	Message string   `json:"message"` //Invalid credentials
	Details []string `json:"details"`
}

//------------------------

type NetellerGetPaymentHandleResp struct {
	Meta struct {
		NumberOfRecords int `json:"numberOfRecords"`
		Limit           int `json:"limit"`
		Page            int `json:"page"`
	} `json:"meta"`
	PaymentHandles []struct {
		Id                 string `json:"id"`
		MerchantRefNum     string `json:"merchantRefNum"`
		PaymentHandleToken string `json:"paymentHandleToken"`
		Status             string `json:"status"`
		PaymentType        string `json:"paymentType"`
		LiveMode           bool   `json:"liveMode"`
		Usage              string `json:"usage"`
		Action             string `json:"action"`
		ExecutionMode      string `json:"executionMode"`
		Amount             int    `json:"amount"`
		CurrencyCode       string `json:"currencyCode"`
		MerchantDescriptor struct {
			DynamicDescriptor string `json:"dynamicDescriptor"`
			Phone             string `json:"phone"`
		} `json:"merchantDescriptor"`
		BillingDetails struct {
			Street1 string `json:"street1"`
			Street2 string `json:"street2"`
			City    string `json:"city"`
			Zip     string `json:"zip"`
			Country string `json:"country"`
		} `json:"billingDetails"`
		CustomerIp        string `json:"customerIp"`
		TimeToLiveSeconds int    `json:"timeToLiveSeconds"`
		GatewayResponse   struct {
			OrderId     string      `json:"orderId"`
			TotalAmount interface{} `json:"totalAmount"` //文档是int,但是实际是string
			Currency    string      `json:"currency"`
			Lang        string      `json:"lang"`
			Status      string      `json:"status"`
			Processor   string      `json:"processor"`
		} `json:"gatewayResponse"`
		ReturnLinks struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"returnLinks"`
		TransactionType string `json:"transactionType"`
		Links           []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Neteller struct {
			ConsumerId       string `json:"consumerId"`
			ConsumerIdLocked bool   `json:"consumerIdLocked"`
		} `json:"neteller"`
	} `json:"paymentHandles"`
}

// ============== callback ======================================================

//https://developer.paysafe.com/en/neteller-api-1/#/#webhooks-events

type NetellerPaymentBackReq struct {
	//https://developer.paysafe.com/en/neteller-api-1/#/#webhooks-events
	EventType     string               `json:"eventType" mapstructure:"eventType"` //枚举: PAYMENT_HANDLE_PAYABLE,PAYMENT_HANDLE_COMPLETED,PAYMENT_HANDLE_FAILED
	AttemptNumber string               `json:"attemptNumber" mapstructure:"attemptNumber"`
	ResourceId    string               `json:"resourceId" mapstructure:"resourceId"`
	EventDate     string               `json:"eventDate" mapstructure:"eventDate"`
	Links         []Link               `json:"links" mapstructure:"links"`
	Mode          string               `json:"mode" mapstructure:"mode"`
	EventName     string               `json:"eventName" mapstructure:"eventName"`
	Payload       PaymentHandlePayload `json:"payload" mapstructure:"payload"`
}

type PaymentHandlePayload struct {
	AccountId      string `json:"accountId" mapstructure:"accountId"`
	ID             string `json:"id" mapstructure:"id"`
	MerchantRefNum string `json:"merchantRefNum" mapstructure:"merchantRefNum"`
	Amount         int    `json:"amount" mapstructure:"amount"`
	CurrencyCode   string `json:"currencyCode" mapstructure:"currencyCode"`
	Status         string `json:"status" mapstructure:"status"`
	PaymentType    string `json:"paymentType" mapstructure:"paymentType"`
	TxnTime        string `json:"txnTime" mapstructure:"txnTime"`
}

// https://developer.paysafe.com/en/neteller-api-1/#/#sample-webhooks-payload
/*
	{
        "payload": {
        "accountId": "1011872745",
        "id": "944b9b19-c07c-4953-98ec-761272aef759",
        "merchantRefNum": "3de43050-a857-4b13-a27d-da35daea253d",
        "amount": 100,
        "currencyCode": "USD",
        "status": "COMPLETED",
        "paymentType": "NETELLER",
        "txnTime": "2020-03-18T11:53:29Z"
        },
        "eventType": "PAYMENT_HANDLE_COMPLETED",
        "attemptNumber": "1",
        "resourceId": "944b9b19-c07c-4953-98ec-761272aef759",
        "eventDate": "2020-03-18T11:53:29Z",
        "links": [
         {
           "href": "https://api.qa.paysafe.com/alternatepayments/v1/accounts/1011872745/paymenthandles/944b9b19-c07c-4953-98ec-761272aef759",
           "rel": "payment_handle"
         }
        ],
        "mode": "live",
        "eventName": "PAYMENT_HANDLE_COMPLETED"
}

*/
