package arrowflightrpc

import (
	"encoding/json"
	"fmt"
	"testing"

	gsf_proto "gokratos/api/gsf/data"

	"github.com/stretchr/testify/assert"
)

var clients = []string{
	"192.168.1.188:30005",
	"192.168.1.188:30015",
	"192.168.1.188:30025",
	"192.168.1.192:30005",
	// "192.168.15.33:50005",
	"192.168.1.188:30035",
	"192.168.1.187:58005",
	"192.168.1.187:59005",
}

func TestXxx(t *testing.T) {

	for _, endpoint := range clients {
		client := NewFlightClient(endpoint)
		err := client.Connect()
		assert.Nil(t, err)
		defer client.Disconnect()

		act, err := client.ActionList()
		assert.Nil(t, err)
		client.DoAction(act)
	}

}

// func TestListFlights(t *testing.T) {
// 	endpoint := "192.168.1.188:30005"
// 	client := NewFlightClient(endpoint)
// 	err := client.Connect()
// 	assert.Nil(t, err)
// 	defer client.Disconnect()

// 	ctx := context.Background()
// 	in := &flight.Criteria{Expression: []byte("")}
// 	listFlightsClient, err := client.client.ListFlights(ctx, in)
// 	assert.Nil(t, err)
// 	flightinfo, err := listFlightsClient.Recv()
// 	fmt.Printf("--> %+v", err)
// 	assert.Nil(t, err)
// }

func TestGetBalances(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_positions",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)
	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.BalanceT, 0)
	for r.Next() {
		bls := []*gsf_proto.BalanceT{}
		record := r.Record()
		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &bls)
		assert.Nil(t, err)
		ret = append(ret, bls...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetPositions(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_positions",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.PositionT, 0)
	for r.Next() {
		poses := []*gsf_proto.PositionT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &poses)
		assert.Nil(t, err)
		ret = append(ret, poses...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetOrders(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_orders",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.OrderT, 0)
	for r.Next() {
		orders := []*gsf_proto.OrderT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &orders)
		assert.Nil(t, err)
		ret = append(ret, orders...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetTrades(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_trades",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.TradeT, 0)
	for r.Next() {
		trades := []*gsf_proto.TradeT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &trades)
		assert.Nil(t, err)
		ret = append(ret, trades...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetGoals(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_trades",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.GoalPositionT, 0)
	for r.Next() {
		goalpos := []*gsf_proto.GoalPositionT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &goalpos)
		assert.Nil(t, err)
		ret = append(ret, goalpos...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetUnfilledOrders(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_unfilled_orders",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.OrderT, 0)
	for r.Next() {
		orders := []*gsf_proto.OrderT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &orders)
		assert.Nil(t, err)
		ret = append(ret, orders...)
	}

	assert.Greater(t, len(ret), 0)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetETs(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_ets",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.EquityTransitT, 0)
	for r.Next() {
		ets := []*gsf_proto.EquityTransitT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &ets)
		assert.Nil(t, err)
		ret = append(ret, ets...)
	}

	assert.Greater(t, len(ret), -1)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

func TestGetFTs(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_fts",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.FundTransitT, 0)
	for r.Next() {
		fts := []*gsf_proto.FundTransitT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &fts)
		assert.Nil(t, err)
		ret = append(ret, fts...)
	}

	assert.Greater(t, len(ret), -1)
	// for _, item := range ret {
	// 	fmt.Printf("%+v\n", item)
	// }
}

// func TestGetInfos(t *testing.T) {
// 	// todo has error
// 	endpoint := "192.168.1.188:30035"
// 	client := NewFlightClient(endpoint)
// 	err := client.Connect()
// 	assert.Nil(t, err)
// 	defer client.Disconnect()

// 	req := map[string]interface{}{
// 		"service_name": "Oms",
// 		"data_name":    "get_infos",
// 	}
// 	buf, err := json.Marshal(req)
// 	assert.Nil(t, err)

// 	r, err := client.DoAction(string(buf))
// 	assert.Nil(t, err)

// 	fmt.Println(string(r))
// }

func TestGetAuMetrics(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)
	err := client.Connect()
	assert.Nil(t, err)
	defer client.Disconnect()

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_au_metrics",
	}
	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	ret := make([]*gsf_proto.AuMetricsT, 0)
	for r.Next() {
		aumetrics := []*gsf_proto.AuMetricsT{}
		record := r.Record()

		jsonBuf, _ := record.MarshalJSON()
		err := json.Unmarshal(jsonBuf, &aumetrics)
		assert.Nil(t, err)
		ret = append(ret, aumetrics...)
	}

	assert.Greater(t, len(ret), -1)
	for _, item := range ret {
		fmt.Printf("%+v\n", item)
	}
}
