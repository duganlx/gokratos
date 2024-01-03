// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ZbOrderT struct {
	EpochLocalTime int64 `json:"epoch_local_time"`
	Symbol string `json:"symbol"`
	Market string `json:"market"`
	StreamId uint32 `json:"stream_id"`
	EpochExchgTime int64 `json:"epoch_exchg_time"`
	Channel uint16 `json:"channel"`
	OrderNo uint64 `json:"order_no"`
	BizIndex uint64 `json:"biz_index"`
	Index uint64 `json:"index"`
	Price float64 `json:"price"`
	Volume int64 `json:"volume"`
	OrderType string `json:"order_type"`
	LocalTime uint32 `json:"local_time"`
	ExchgTime uint32 `json:"exchg_time"`
	TradeDate int64 `json:"trade_date"`
	Key int32 `json:"key"`
	EamCode int64 `json:"eam_code"`
	Tup float64 `json:"tup"`
}

func (t *ZbOrderT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	symbolOffset := flatbuffers.UOffsetT(0)
	if t.Symbol != "" {
		symbolOffset = builder.CreateString(t.Symbol)
	}
	marketOffset := flatbuffers.UOffsetT(0)
	if t.Market != "" {
		marketOffset = builder.CreateString(t.Market)
	}
	orderTypeOffset := flatbuffers.UOffsetT(0)
	if t.OrderType != "" {
		orderTypeOffset = builder.CreateString(t.OrderType)
	}
	ZbOrderStart(builder)
	ZbOrderAddEpochLocalTime(builder, t.EpochLocalTime)
	ZbOrderAddSymbol(builder, symbolOffset)
	ZbOrderAddMarket(builder, marketOffset)
	ZbOrderAddStreamId(builder, t.StreamId)
	ZbOrderAddEpochExchgTime(builder, t.EpochExchgTime)
	ZbOrderAddChannel(builder, t.Channel)
	ZbOrderAddOrderNo(builder, t.OrderNo)
	ZbOrderAddBizIndex(builder, t.BizIndex)
	ZbOrderAddIndex(builder, t.Index)
	ZbOrderAddPrice(builder, t.Price)
	ZbOrderAddVolume(builder, t.Volume)
	ZbOrderAddOrderType(builder, orderTypeOffset)
	ZbOrderAddLocalTime(builder, t.LocalTime)
	ZbOrderAddExchgTime(builder, t.ExchgTime)
	ZbOrderAddTradeDate(builder, t.TradeDate)
	ZbOrderAddKey(builder, t.Key)
	ZbOrderAddEamCode(builder, t.EamCode)
	ZbOrderAddTup(builder, t.Tup)
	return ZbOrderEnd(builder)
}

func (rcv *ZbOrder) UnPackTo(t *ZbOrderT) {
	t.EpochLocalTime = rcv.EpochLocalTime()
	t.Symbol = string(rcv.Symbol())
	t.Market = string(rcv.Market())
	t.StreamId = rcv.StreamId()
	t.EpochExchgTime = rcv.EpochExchgTime()
	t.Channel = rcv.Channel()
	t.OrderNo = rcv.OrderNo()
	t.BizIndex = rcv.BizIndex()
	t.Index = rcv.Index()
	t.Price = rcv.Price()
	t.Volume = rcv.Volume()
	t.OrderType = string(rcv.OrderType())
	t.LocalTime = rcv.LocalTime()
	t.ExchgTime = rcv.ExchgTime()
	t.TradeDate = rcv.TradeDate()
	t.Key = rcv.Key()
	t.EamCode = rcv.EamCode()
	t.Tup = rcv.Tup()
}

func (rcv *ZbOrder) UnPack() *ZbOrderT {
	if rcv == nil {
		return nil
	}
	t := &ZbOrderT{}
	rcv.UnPackTo(t)
	return t
}

type ZbOrder struct {
	_tab flatbuffers.Table
}

func GetRootAsZbOrder(buf []byte, offset flatbuffers.UOffsetT) *ZbOrder {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ZbOrder{}
	x.Init(buf, n+offset)
	return x
}

func FinishZbOrderBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsZbOrder(buf []byte, offset flatbuffers.UOffsetT) *ZbOrder {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ZbOrder{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedZbOrderBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ZbOrder) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ZbOrder) Table() flatbuffers.Table {
	return rcv._tab
}

/// 本地接收日期时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) EpochLocalTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// 本地接收日期时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) MutateEpochLocalTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

/// 证券代码
func (rcv *ZbOrder) Symbol() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// 证券代码
///市场
///SZ,SH,HK,CF,BJ,O,N
func (rcv *ZbOrder) Market() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///市场
///SZ,SH,HK,CF,BJ,O,N
/// 行情类别，新增
func (rcv *ZbOrder) StreamId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// 行情类别，新增
func (rcv *ZbOrder) MutateStreamId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

/// 交易所行情日期时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) EpochExchgTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// 交易所行情日期时间，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) MutateEpochExchgTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(12, n)
}

/// 通道 ，新增
func (rcv *ZbOrder) Channel() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

/// 通道 ，新增
func (rcv *ZbOrder) MutateChannel(n uint16) bool {
	return rcv._tab.MutateUint16Slot(14, n)
}

/// 订单编号号
/// 1. 深交所原始行情中该字段为index, GSF2.0中会复制index到order_no
/// 2. 上交所原始行情中该字段为新增、删除订单时用以标识订单的唯一编号[order id]
func (rcv *ZbOrder) OrderNo() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// 订单编号号
/// 1. 深交所原始行情中该字段为index, GSF2.0中会复制index到order_no
/// 2. 上交所原始行情中该字段为新增、删除订单时用以标识订单的唯一编号[order id]
func (rcv *ZbOrder) MutateOrderNo(n uint64) bool {
	return rcv._tab.MutateUint64Slot(16, n)
}

/// 业务序列号(沪市特有)
/// 每个通道（Channel）内逐笔成交数据与逐笔委托数据统一排序生成业务序列号（BizIndex），并从 1 开始递增。新增
func (rcv *ZbOrder) BizIndex() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// 业务序列号(沪市特有)
/// 每个通道（Channel）内逐笔成交数据与逐笔委托数据统一排序生成业务序列号（BizIndex），并从 1 开始递增。新增
func (rcv *ZbOrder) MutateBizIndex(n uint64) bool {
	return rcv._tab.MutateUint64Slot(18, n)
}

/// 消息编号
/// 1. 深交所原始行情中index作为原始订单编号, GSF2.0中会复制到order_no
/// 2. 上交所原始行情中index为消息编号, 原始订单编号为order_no
func (rcv *ZbOrder) Index() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// 消息编号
/// 1. 深交所原始行情中index作为原始订单编号, GSF2.0中会复制到order_no
/// 2. 上交所原始行情中index为消息编号, 原始订单编号为order_no
func (rcv *ZbOrder) MutateIndex(n uint64) bool {
	return rcv._tab.MutateUint64Slot(20, n)
}

/// 委托价格
func (rcv *ZbOrder) Price() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// 委托价格
func (rcv *ZbOrder) MutatePrice(n float64) bool {
	return rcv._tab.MutateFloat64Slot(22, n)
}

/// 委托量（注意沪市剩余量的概念）
func (rcv *ZbOrder) Volume() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// 委托量（注意沪市剩余量的概念）
func (rcv *ZbOrder) MutateVolume(n int64) bool {
	return rcv._tab.MutateInt64Slot(24, n)
}

/// (0 买卖方向:B=买,S=卖,C=撤,G=借入,F=出借);
/// (1 深市：委托类别'0'=限价, '1'=市价，'U'=本方最优.   沪市：订单类型'A'=增加, 'D'=删除);
/// (2-3 保留)
func (rcv *ZbOrder) OrderType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// (0 买卖方向:B=买,S=卖,C=撤,G=借入,F=出借);
/// (1 深市：委托类别'0'=限价, '1'=市价，'U'=本方最优.   沪市：订单类型'A'=增加, 'D'=删除);
/// (2-3 保留)
/// 本地接收时间(整数形式)，HHMMSSmmm 样例：93000500或者153000000
func (rcv *ZbOrder) LocalTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// 本地接收时间(整数形式)，HHMMSSmmm 样例：93000500或者153000000
func (rcv *ZbOrder) MutateLocalTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(28, n)
}

/// 交易所时间(整数形式)，HHMMSSmmm 样例：93000500或者153000000
func (rcv *ZbOrder) ExchgTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// 交易所时间(整数形式)，HHMMSSmmm 样例：93000500或者153000000
func (rcv *ZbOrder) MutateExchgTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(30, n)
}

/// 交易日，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) TradeDate() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// 交易日，Epoch时间(13位数字，UTC时区)，样例：1676017139000
func (rcv *ZbOrder) MutateTradeDate(n int64) bool {
	return rcv._tab.MutateInt64Slot(32, n)
}

/// code_id的hash值，股票为code_id的数字，期货为code_id字符串的hash值
func (rcv *ZbOrder) Key() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// code_id的hash值，股票为code_id的数字，期货为code_id字符串的hash值
func (rcv *ZbOrder) MutateKey(n int32) bool {
	return rcv._tab.MutateInt32Slot(34, n)
}

/// eam内码
func (rcv *ZbOrder) EamCode() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// eam内码
func (rcv *ZbOrder) MutateEamCode(n int64) bool {
	return rcv._tab.MutateInt64Slot(36, n)
}

/// 奔放最优的价格【回验撮合专用】
func (rcv *ZbOrder) Tup() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return -1.0
}

/// 奔放最优的价格【回验撮合专用】
func (rcv *ZbOrder) MutateTup(n float64) bool {
	return rcv._tab.MutateFloat64Slot(38, n)
}

func ZbOrderStart(builder *flatbuffers.Builder) {
	builder.StartObject(18)
}
func ZbOrderAddEpochLocalTime(builder *flatbuffers.Builder, epochLocalTime int64) {
	builder.PrependInt64Slot(0, epochLocalTime, 0)
}
func ZbOrderAddSymbol(builder *flatbuffers.Builder, symbol flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(symbol), 0)
}
func ZbOrderAddMarket(builder *flatbuffers.Builder, market flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(market), 0)
}
func ZbOrderAddStreamId(builder *flatbuffers.Builder, streamId uint32) {
	builder.PrependUint32Slot(3, streamId, 0)
}
func ZbOrderAddEpochExchgTime(builder *flatbuffers.Builder, epochExchgTime int64) {
	builder.PrependInt64Slot(4, epochExchgTime, 0)
}
func ZbOrderAddChannel(builder *flatbuffers.Builder, channel uint16) {
	builder.PrependUint16Slot(5, channel, 0)
}
func ZbOrderAddOrderNo(builder *flatbuffers.Builder, orderNo uint64) {
	builder.PrependUint64Slot(6, orderNo, 0)
}
func ZbOrderAddBizIndex(builder *flatbuffers.Builder, bizIndex uint64) {
	builder.PrependUint64Slot(7, bizIndex, 0)
}
func ZbOrderAddIndex(builder *flatbuffers.Builder, index uint64) {
	builder.PrependUint64Slot(8, index, 0)
}
func ZbOrderAddPrice(builder *flatbuffers.Builder, price float64) {
	builder.PrependFloat64Slot(9, price, 0.0)
}
func ZbOrderAddVolume(builder *flatbuffers.Builder, volume int64) {
	builder.PrependInt64Slot(10, volume, 0)
}
func ZbOrderAddOrderType(builder *flatbuffers.Builder, orderType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(orderType), 0)
}
func ZbOrderAddLocalTime(builder *flatbuffers.Builder, localTime uint32) {
	builder.PrependUint32Slot(12, localTime, 0)
}
func ZbOrderAddExchgTime(builder *flatbuffers.Builder, exchgTime uint32) {
	builder.PrependUint32Slot(13, exchgTime, 0)
}
func ZbOrderAddTradeDate(builder *flatbuffers.Builder, tradeDate int64) {
	builder.PrependInt64Slot(14, tradeDate, 0)
}
func ZbOrderAddKey(builder *flatbuffers.Builder, key int32) {
	builder.PrependInt32Slot(15, key, -1)
}
func ZbOrderAddEamCode(builder *flatbuffers.Builder, eamCode int64) {
	builder.PrependInt64Slot(16, eamCode, 0)
}
func ZbOrderAddTup(builder *flatbuffers.Builder, tup float64) {
	builder.PrependFloat64Slot(17, tup, -1.0)
}
func ZbOrderEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}