// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SendOrderCommandT struct {
	Time int64 `json:"time"`
	Order *OrderT `json:"order"`
}

func (t *SendOrderCommandT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	orderOffset := t.Order.Pack(builder)
	SendOrderCommandStart(builder)
	SendOrderCommandAddTime(builder, t.Time)
	SendOrderCommandAddOrder(builder, orderOffset)
	return SendOrderCommandEnd(builder)
}

func (rcv *SendOrderCommand) UnPackTo(t *SendOrderCommandT) {
	t.Time = rcv.Time()
	t.Order = rcv.Order(nil).UnPack()
}

func (rcv *SendOrderCommand) UnPack() *SendOrderCommandT {
	if rcv == nil {
		return nil
	}
	t := &SendOrderCommandT{}
	rcv.UnPackTo(t)
	return t
}

type SendOrderCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsSendOrderCommand(buf []byte, offset flatbuffers.UOffsetT) *SendOrderCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SendOrderCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishSendOrderCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSendOrderCommand(buf []byte, offset flatbuffers.UOffsetT) *SendOrderCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SendOrderCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSendOrderCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SendOrderCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SendOrderCommand) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SendOrderCommand) Time() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SendOrderCommand) MutateTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *SendOrderCommand) Order(obj *Order) *Order {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Order)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SendOrderCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SendOrderCommandAddTime(builder *flatbuffers.Builder, time int64) {
	builder.PrependInt64Slot(0, time, 0)
}
func SendOrderCommandAddOrder(builder *flatbuffers.Builder, order flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(order), 0)
}
func SendOrderCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
