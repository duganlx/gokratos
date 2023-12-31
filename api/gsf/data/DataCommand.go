// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DataCommandT struct {
	Time int64 `json:"time"`
	Topic string `json:"topic"`
	Data string `json:"data"`
	Mode int16 `json:"mode"`
}

func (t *DataCommandT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	topicOffset := flatbuffers.UOffsetT(0)
	if t.Topic != "" {
		topicOffset = builder.CreateString(t.Topic)
	}
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != "" {
		dataOffset = builder.CreateString(t.Data)
	}
	DataCommandStart(builder)
	DataCommandAddTime(builder, t.Time)
	DataCommandAddTopic(builder, topicOffset)
	DataCommandAddData(builder, dataOffset)
	DataCommandAddMode(builder, t.Mode)
	return DataCommandEnd(builder)
}

func (rcv *DataCommand) UnPackTo(t *DataCommandT) {
	t.Time = rcv.Time()
	t.Topic = string(rcv.Topic())
	t.Data = string(rcv.Data())
	t.Mode = rcv.Mode()
}

func (rcv *DataCommand) UnPack() *DataCommandT {
	if rcv == nil {
		return nil
	}
	t := &DataCommandT{}
	rcv.UnPackTo(t)
	return t
}

type DataCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsDataCommand(buf []byte, offset flatbuffers.UOffsetT) *DataCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishDataCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDataCommand(buf []byte, offset flatbuffers.UOffsetT) *DataCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDataCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DataCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataCommand) Table() flatbuffers.Table {
	return rcv._tab
}

/// 时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *DataCommand) Time() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// 时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *DataCommand) MutateTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *DataCommand) Topic() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DataCommand) Data() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// 时间模式，0 - History, 1 - Realtime, 2 - Both
func (rcv *DataCommand) Mode() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

/// 时间模式，0 - History, 1 - Realtime, 2 - Both
func (rcv *DataCommand) MutateMode(n int16) bool {
	return rcv._tab.MutateInt16Slot(10, n)
}

func DataCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DataCommandAddTime(builder *flatbuffers.Builder, time int64) {
	builder.PrependInt64Slot(0, time, 0)
}
func DataCommandAddTopic(builder *flatbuffers.Builder, topic flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(topic), 0)
}
func DataCommandAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(data), 0)
}
func DataCommandAddMode(builder *flatbuffers.Builder, mode int16) {
	builder.PrependInt16Slot(3, mode, 0)
}
func DataCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
