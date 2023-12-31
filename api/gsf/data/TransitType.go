// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import "strconv"

type TransitType int16

const (
	/// 支出
	TransitTypeCOST     TransitType = -1
	/// 默认
	TransitTypeDEFAULT  TransitType = 0
	/// 收入
	TransitTypeINCOME   TransitType = 1
	/// 账户间划转
	TransitTypeTRANSFER TransitType = 2
	/// 手动冲销
	TransitTypeMANUAL   TransitType = 3
	/// 融资
	TransitTypeLOAN     TransitType = 4
)

var EnumNamesTransitType = map[TransitType]string{
	TransitTypeCOST:     "COST",
	TransitTypeDEFAULT:  "DEFAULT",
	TransitTypeINCOME:   "INCOME",
	TransitTypeTRANSFER: "TRANSFER",
	TransitTypeMANUAL:   "MANUAL",
	TransitTypeLOAN:     "LOAN",
}

var EnumValuesTransitType = map[string]TransitType{
	"COST":     TransitTypeCOST,
	"DEFAULT":  TransitTypeDEFAULT,
	"INCOME":   TransitTypeINCOME,
	"TRANSFER": TransitTypeTRANSFER,
	"MANUAL":   TransitTypeMANUAL,
	"LOAN":     TransitTypeLOAN,
}

func (v TransitType) String() string {
	if s, ok := EnumNamesTransitType[v]; ok {
		return s
	}
	return "TransitType(" + strconv.FormatInt(int64(v), 10) + ")"
}
