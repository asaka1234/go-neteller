package go_neteller

type NetellerInitParams struct {
	MerchantId  string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"`    // merchantId
	MerchantKey string `json:"merchantKey" mapstructure:"merchantKey" config:"merchantKey"` // accessKey

	CreatePaymentHandleUrl string `json:"createPaymentHandleUrl" mapstructure:"createPaymentHandleUrl" config:"createPaymentHandleUrl"`
	PaymentBackUrl         string `json:"paymentBackUrl" mapstructure:"paymentBackUrl" config:"paymentBackUrl"` //回调地址
}

// ----------pre generate-------------------------

type NetellerPaymentHandleReq struct {
	MerchantRefNum  string         `json:"merchantRefNum" mapstructure:"merchantRefNum"`
	TransactionType string         `json:"transactionType" mapstructure:"transactionType"` //枚举: PAYMENT (付款给商户)， STANDALONE_CREDIT (商户付给User)
	Amount          int            `json:"amount" mapstructure:"amount"`                   //这里需要做单位转换,用的是法币的最小单位. 比如1.2美元，这里传的是120美分
	CurrencyCode    string         `json:"currencyCode" mapstructure:"currencyCode"`       //币种
	CustomerIp      string         `json:"customerIp" mapstructure:"customerIp"`           //客户ip
	Neteller        NetellerDetail `json:"neteller" mapstructure:"neteller"`
	//sdk搞定
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

type PaymentRequest struct {
	TransactionType string         `json:"transactionType"` //指定是那种请求, 枚举: PAYMENT,STANDALONE_CREDIT .决策是充值/提现
	MerchantRefNum  string         `json:"merchantRefNum"`  //随机唯一num, This is the merchant reference number created by you
	PaymentType     string         `json:"paymentType"`     //付款类型，e.g.NETELLER
	Amount          int            `json:"amount"`          //这里做了单位转换,用的是法币的最小单位. 比如1.2美元，这里传的是120美分
	CurrencyCode    string         `json:"currencyCode"`    //币种类型
	ReturnLinks     []ReturnLink   `json:"returnLinks"`     //设置回调url
	Neteller        NetellerDetail `json:"neteller"`
	//option
	CustomerIp     string          `json:"customerIp"`     //customer's IP address
	BillingDetails *BillingDetails `json:"billingDetails"` // customer's billing details
}

//-----------------------------------------------------------

type NetellerPaymentHandleResp struct {
	//error
	Error ErrorInfo `json:"error"` //psp三方的订单号
	//succeed
	ID                 string `json:"id" mapstructure:"id"`                                 //代表这个session
	PaymentHandleToken string `json:"paymentHandleToken" mapstructure:"paymentHandleToken"` //可以认为是psp的订单号
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
	OrderId     string `json:"orderId" mapstructure:"orderId"`         //psp的订单号
	TotalAmount int    `json:"totalAmount" mapstructure:"totalAmount"` //The total amount due for this order, including all items, fees, taxes
	Currency    string `json:"currency" mapstructure:"currency"`
	Status      string `json:"status" mapstructure:"status"` //枚举: pending,cancelled,failed (The order was not paid).,paid,expired (The order had expired. Default expiry time is 15 mins).
	Lang        string `json:"lang" mapstructure:"lang"`     //en_US
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
