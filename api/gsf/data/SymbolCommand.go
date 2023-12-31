// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SymbolCommandT struct {
	TradeDate string `json:"trade_date"`
	Symbols []*SymbolT `json:"symbols"`
}

func (t *SymbolCommandT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	tradeDateOffset := flatbuffers.UOffsetT(0)
	if t.TradeDate != "" {
		tradeDateOffset = builder.CreateString(t.TradeDate)
	}
	symbolsOffset := flatbuffers.UOffsetT(0)
	if t.Symbols != nil {
		symbolsLength := len(t.Symbols)
		symbolsOffsets := make([]flatbuffers.UOffsetT, symbolsLength)
		for j := 0; j < symbolsLength; j++ {
			symbolsOffsets[j] = t.Symbols[j].Pack(builder)
		}
		SymbolCommandStartSymbolsVector(builder, symbolsLength)
		for j := symbolsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(symbolsOffsets[j])
		}
		symbolsOffset = builder.EndVector(symbolsLength)
	}
	SymbolCommandStart(builder)
	SymbolCommandAddTradeDate(builder, tradeDateOffset)
	SymbolCommandAddSymbols(builder, symbolsOffset)
	return SymbolCommandEnd(builder)
}

func (rcv *SymbolCommand) UnPackTo(t *SymbolCommandT) {
	t.TradeDate = string(rcv.TradeDate())
	symbolsLength := rcv.SymbolsLength()
	t.Symbols = make([]*SymbolT, symbolsLength)
	for j := 0; j < symbolsLength; j++ {
		x := Symbol{}
		rcv.Symbols(&x, j)
		t.Symbols[j] = x.UnPack()
	}
}

func (rcv *SymbolCommand) UnPack() *SymbolCommandT {
	if rcv == nil {
		return nil
	}
	t := &SymbolCommandT{}
	rcv.UnPackTo(t)
	return t
}

type SymbolCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsSymbolCommand(buf []byte, offset flatbuffers.UOffsetT) *SymbolCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SymbolCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishSymbolCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSymbolCommand(buf []byte, offset flatbuffers.UOffsetT) *SymbolCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SymbolCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSymbolCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SymbolCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SymbolCommand) Table() flatbuffers.Table {
	return rcv._tab
}

/// 日期，格式：YYYYMMDD
func (rcv *SymbolCommand) TradeDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// 日期，格式：YYYYMMDD
func (rcv *SymbolCommand) Symbols(obj *Symbol, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SymbolCommand) SymbolsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SymbolCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SymbolCommandAddTradeDate(builder *flatbuffers.Builder, tradeDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(tradeDate), 0)
}
func SymbolCommandAddSymbols(builder *flatbuffers.Builder, symbols flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(symbols), 0)
}
func SymbolCommandStartSymbolsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SymbolCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
