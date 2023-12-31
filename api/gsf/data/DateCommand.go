// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DateCommandT struct {
	TradeDate string `json:"trade_date"`
}

func (t *DateCommandT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	tradeDateOffset := flatbuffers.UOffsetT(0)
	if t.TradeDate != "" {
		tradeDateOffset = builder.CreateString(t.TradeDate)
	}
	DateCommandStart(builder)
	DateCommandAddTradeDate(builder, tradeDateOffset)
	return DateCommandEnd(builder)
}

func (rcv *DateCommand) UnPackTo(t *DateCommandT) {
	t.TradeDate = string(rcv.TradeDate())
}

func (rcv *DateCommand) UnPack() *DateCommandT {
	if rcv == nil {
		return nil
	}
	t := &DateCommandT{}
	rcv.UnPackTo(t)
	return t
}

type DateCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsDateCommand(buf []byte, offset flatbuffers.UOffsetT) *DateCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DateCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishDateCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDateCommand(buf []byte, offset flatbuffers.UOffsetT) *DateCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DateCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDateCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DateCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DateCommand) Table() flatbuffers.Table {
	return rcv._tab
}

/// 日期，格式：YYYYMMDD
func (rcv *DateCommand) TradeDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// 日期，格式：YYYYMMDD
func DateCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DateCommandAddTradeDate(builder *flatbuffers.Builder, tradeDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(tradeDate), 0)
}
func DateCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
