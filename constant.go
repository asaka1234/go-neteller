package go_neteller

import "strings"

// OutType 出款类型枚举
type OutType int

const (
	Deposit  OutType = 1 // 入金
	Withdraw OutType = 2 // 出金
)

// String 返回枚举的字符串表示
func (o OutType) String() string {
	switch o {
	case Deposit:
		return "deposit"
	case Withdraw:
		return "withdraw"
	default:
		return "unknown"
	}
}

// Desc 返回枚举的描述
func (o OutType) Desc() string {
	switch o {
	case Deposit:
		return "入金"
	case Withdraw:
		return "出金"
	default:
		return "未知类型"
	}
}

// Eq 比较是否相等
func (o OutType) Eq(value int) bool {
	return int(o) == value
}

// Name 返回枚举名称(同String)
func (o OutType) Name() string {
	return o.String()
}

// Value 返回枚举值
func (o OutType) Value() int {
	return int(o)
}

// OutTypeInfo 枚举信息结构体
type OutTypeInfo struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// GetInfo 获取枚举的完整信息
func (o OutType) GetInfo() OutTypeInfo {
	return OutTypeInfo{
		Code: o.Value(),
		Name: o.Name(),
		Desc: o.Desc(),
	}
}

// GetAllOutTypes 获取所有出款类型
func GetAllOutTypes() []OutTypeInfo {
	return []OutTypeInfo{
		Deposit.GetInfo(),
		Withdraw.GetInfo(),
	}
}

//============================================

type NetellerEventTypeEnum struct {
	code string
	name string
}

func (e NetellerEventTypeEnum) GetName() string {
	return e.name
}

func (e NetellerEventTypeEnum) GetValue() string {
	return e.code
}

func (e NetellerEventTypeEnum) Eq(value string) bool {
	return e.code == value
}

var (
	PaymentHandlePayable   = NetellerEventTypeEnum{"PAYMENT_HANDLE_PAYABLE", "PAYMENT_HANDLE_PAYABLE"}
	PaymentHandleCompleted = NetellerEventTypeEnum{"PAYMENT_HANDLE_COMPLETED", "PAYMENT_HANDLE_COMPLETED"}
	PaymentHandleFailed    = NetellerEventTypeEnum{"PAYMENT_HANDLE_FAILED", "PAYMENT_HANDLE_FAILED"}
	PaymentHandleExpired   = NetellerEventTypeEnum{"PAYMENT_HANDLE_EXPIRED", "PAYMENT_HANDLE_EXPIRED"}

	PaymentCompleted = NetellerEventTypeEnum{"PAYMENT_COMPLETED", "PAYMENT_COMPLETED"}
	PaymentHeld      = NetellerEventTypeEnum{"PAYMENT_HELD", "PAYMENT_HELD"}
	PaymentFailed    = NetellerEventTypeEnum{"PAYMENT_FAILED", "PAYMENT_FAILED"}

	SaCreditCompleted = NetellerEventTypeEnum{"SA_CREDIT_COMPLETED", "SA_CREDIT_COMPLETED"}
	SaCreditHeld      = NetellerEventTypeEnum{"SA_CREDIT_HELD", "SA_CREDIT_HELD"}
	SaCreditFailed    = NetellerEventTypeEnum{"SA_CREDIT_FAILED", "SA_CREDIT_FAILED"}
	SaCreditCancelled = NetellerEventTypeEnum{"SA_CREDIT_CANCELLED", "SA_CREDIT_CANCELLED"}
)

// StringToNetellerEventTypeEnum converts a string to the corresponding NetellerEventTypeEnum
func StringToNetellerEventTypeEnum(s string) (NetellerEventTypeEnum, bool) {
	s = strings.ToUpper(s)
	switch s {
	case "PAYMENT_HANDLE_PAYABLE":
		return PaymentHandlePayable, true
	case "PAYMENT_HANDLE_COMPLETED":
		return PaymentHandleCompleted, true
	case "PAYMENT_HANDLE_FAILED":
		return PaymentHandleFailed, true
	case "PAYMENT_HANDLE_EXPIRED":
		return PaymentHandleExpired, true
	case "PAYMENT_COMPLETED":
		return PaymentCompleted, true
	case "PAYMENT_HELD":
		return PaymentHeld, true
	case "PAYMENT_FAILED":
		return PaymentFailed, true
	case "SA_CREDIT_COMPLETED":
		return SaCreditCompleted, true
	case "SA_CREDIT_HELD":
		return SaCreditHeld, true
	case "SA_CREDIT_FAILED":
		return SaCreditFailed, true
	case "SA_CREDIT_CANCELLED":
		return SaCreditCancelled, true
	default:
		return NetellerEventTypeEnum{}, false
	}
}
