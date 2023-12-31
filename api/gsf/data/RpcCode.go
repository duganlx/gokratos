// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import "strconv"

type RpcCode int16

const (
	RpcCodeNONE    RpcCode = 0
	RpcCodeSUCCESS RpcCode = 1
	RpcCodeFAILED  RpcCode = 2
	RpcCodeNO_AUTH RpcCode = 3
)

var EnumNamesRpcCode = map[RpcCode]string{
	RpcCodeNONE:    "NONE",
	RpcCodeSUCCESS: "SUCCESS",
	RpcCodeFAILED:  "FAILED",
	RpcCodeNO_AUTH: "NO_AUTH",
}

var EnumValuesRpcCode = map[string]RpcCode{
	"NONE":    RpcCodeNONE,
	"SUCCESS": RpcCodeSUCCESS,
	"FAILED":  RpcCodeFAILED,
	"NO_AUTH": RpcCodeNO_AUTH,
}

func (v RpcCode) String() string {
	if s, ok := EnumNamesRpcCode[v]; ok {
		return s
	}
	return "RpcCode(" + strconv.FormatInt(int64(v), 10) + ")"
}
