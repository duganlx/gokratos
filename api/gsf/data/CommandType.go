// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import "strconv"

type CommandType uint16

const (
	CommandTypeOnBeforeTrading CommandType = 0
	CommandTypeOnMarketClose   CommandType = 1
	CommandTypeOnAuctionStart  CommandType = 2
	CommandTypeOnAuctionEnd    CommandType = 3
	CommandTypeOnAfterTrading  CommandType = 4
	CommandTypeOnBar           CommandType = 5
	CommandTypeOnData          CommandType = 6
	CommandTypeOnInit          CommandType = 7
	CommandTypeOnDestroy       CommandType = 8
	CommandTypeSetData         CommandType = 9
	CommandTypeOnNotify        CommandType = 10
)

var EnumNamesCommandType = map[CommandType]string{
	CommandTypeOnBeforeTrading: "OnBeforeTrading",
	CommandTypeOnMarketClose:   "OnMarketClose",
	CommandTypeOnAuctionStart:  "OnAuctionStart",
	CommandTypeOnAuctionEnd:    "OnAuctionEnd",
	CommandTypeOnAfterTrading:  "OnAfterTrading",
	CommandTypeOnBar:           "OnBar",
	CommandTypeOnData:          "OnData",
	CommandTypeOnInit:          "OnInit",
	CommandTypeOnDestroy:       "OnDestroy",
	CommandTypeSetData:         "SetData",
	CommandTypeOnNotify:        "OnNotify",
}

var EnumValuesCommandType = map[string]CommandType{
	"OnBeforeTrading": CommandTypeOnBeforeTrading,
	"OnMarketClose":   CommandTypeOnMarketClose,
	"OnAuctionStart":  CommandTypeOnAuctionStart,
	"OnAuctionEnd":    CommandTypeOnAuctionEnd,
	"OnAfterTrading":  CommandTypeOnAfterTrading,
	"OnBar":           CommandTypeOnBar,
	"OnData":          CommandTypeOnData,
	"OnInit":          CommandTypeOnInit,
	"OnDestroy":       CommandTypeOnDestroy,
	"SetData":         CommandTypeSetData,
	"OnNotify":        CommandTypeOnNotify,
}

func (v CommandType) String() string {
	if s, ok := EnumNamesCommandType[v]; ok {
		return s
	}
	return "CommandType(" + strconv.FormatInt(int64(v), 10) + ")"
}