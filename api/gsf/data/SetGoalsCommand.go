// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SetGoalsCommandT struct {
	Time int64 `json:"time"`
	Goals []*GoalPositionT `json:"goals"`
}

func (t *SetGoalsCommandT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	goalsOffset := flatbuffers.UOffsetT(0)
	if t.Goals != nil {
		goalsLength := len(t.Goals)
		goalsOffsets := make([]flatbuffers.UOffsetT, goalsLength)
		for j := 0; j < goalsLength; j++ {
			goalsOffsets[j] = t.Goals[j].Pack(builder)
		}
		SetGoalsCommandStartGoalsVector(builder, goalsLength)
		for j := goalsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(goalsOffsets[j])
		}
		goalsOffset = builder.EndVector(goalsLength)
	}
	SetGoalsCommandStart(builder)
	SetGoalsCommandAddTime(builder, t.Time)
	SetGoalsCommandAddGoals(builder, goalsOffset)
	return SetGoalsCommandEnd(builder)
}

func (rcv *SetGoalsCommand) UnPackTo(t *SetGoalsCommandT) {
	t.Time = rcv.Time()
	goalsLength := rcv.GoalsLength()
	t.Goals = make([]*GoalPositionT, goalsLength)
	for j := 0; j < goalsLength; j++ {
		x := GoalPosition{}
		rcv.Goals(&x, j)
		t.Goals[j] = x.UnPack()
	}
}

func (rcv *SetGoalsCommand) UnPack() *SetGoalsCommandT {
	if rcv == nil {
		return nil
	}
	t := &SetGoalsCommandT{}
	rcv.UnPackTo(t)
	return t
}

type SetGoalsCommand struct {
	_tab flatbuffers.Table
}

func GetRootAsSetGoalsCommand(buf []byte, offset flatbuffers.UOffsetT) *SetGoalsCommand {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SetGoalsCommand{}
	x.Init(buf, n+offset)
	return x
}

func FinishSetGoalsCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSetGoalsCommand(buf []byte, offset flatbuffers.UOffsetT) *SetGoalsCommand {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SetGoalsCommand{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSetGoalsCommandBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SetGoalsCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SetGoalsCommand) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SetGoalsCommand) Time() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SetGoalsCommand) MutateTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *SetGoalsCommand) Goals(obj *GoalPosition, j int) bool {
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

func (rcv *SetGoalsCommand) GoalsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SetGoalsCommandStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SetGoalsCommandAddTime(builder *flatbuffers.Builder, time int64) {
	builder.PrependInt64Slot(0, time, 0)
}
func SetGoalsCommandAddGoals(builder *flatbuffers.Builder, goals flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(goals), 0)
}
func SetGoalsCommandStartGoalsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SetGoalsCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
