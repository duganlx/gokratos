package arrowflightrpc

import (
	"encoding/json"
	"fmt"
	"testing"

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

func TestGetPositions(t *testing.T) {
	endpoint := "192.168.1.188:30005"
	client := NewFlightClient(endpoint)

	req := map[string]interface{}{
		"service_name": "Oms",
		"data_name":    "get_positions",
	}

	buf, err := json.Marshal(req)
	assert.Nil(t, err)

	r, err := client.DoGet(buf)
	assert.Nil(t, err)

	fmt.Println(r)
}
