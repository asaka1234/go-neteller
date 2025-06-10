package go_neteller

import "strings"

type NetellerEventType string

const (
	PaymentHandlePayable   NetellerEventType = "PAYMENT_HANDLE_PAYABLE"
	PaymentHandleCompleted NetellerEventType = "PAYMENT_HANDLE_COMPLETED"
	PaymentHandleFailed    NetellerEventType = "PAYMENT_HANDLE_FAILED"
	PaymentHandleExpired   NetellerEventType = "PAYMENT_HANDLE_EXPIRED"

	PaymentCompleted NetellerEventType = "PAYMENT_COMPLETED"
	PaymentHeld      NetellerEventType = "PAYMENT_HELD"
	PaymentFailed    NetellerEventType = "PAYMENT_FAILED"

	SaCreditCompleted NetellerEventType = "SA_CREDIT_COMPLETED"
	SaCreditHeld      NetellerEventType = "SA_CREDIT_HELD"
	SaCreditFailed    NetellerEventType = "SA_CREDIT_FAILED"
	SaCreditCancelled NetellerEventType = "SA_CREDIT_CANCELLED"
)

// GetName returns the name of the event type (same as value in this case)
func (n NetellerEventType) GetName() string {
	return string(n)
}

// GetValue returns the value of the event type
func (n NetellerEventType) GetValue() string {
	return string(n)
}

// Eq checks if the event type equals the given value (case-insensitive)
func (n NetellerEventType) Eq(value string) bool {
	return strings.EqualFold(string(n), value)
}

// String implements the Stringer interface
func (n NetellerEventType) String() string {
	return string(n)
}
