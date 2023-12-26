package arrowflightrpc

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow/flight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FlightClient struct {
	addr   string
	client flight.Client
}

func NewFlightClient(addr string) *FlightClient {
	return &FlightClient{
		addr: addr,
	}
}

func (c *FlightClient) Connect() error {
	client, err := flight.NewClientWithMiddleware(c.addr, nil, nil, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *FlightClient) Disconnect() error {
	return c.client.Close()
}

func (c *FlightClient) DoGet(buf []byte) (*flight.Reader, error) {
	ctx := context.Background()
	tick := flight.Ticket{
		Ticket: buf,
	}
	fdata, err := c.client.DoGet(ctx, &tick)
	if err != nil {
		return nil, err
	}

	return flight.NewRecordReader(fdata)
}

func (c *FlightClient) DoAction(t string) ([]byte, error) {
	ctx := context.Background()
	act := flight.Action{
		Type: t,
	}
	fdata, err := c.client.DoAction(ctx, &act)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result, err := fdata.Recv()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("do action: %s\n", result.String())

	return result.Body, nil
}

func (c *FlightClient) ActionList() (string, error) {
	ctx := context.Background()
	result, err := c.client.ListActions(ctx, &flight.Empty{})
	if err != nil {
		fmt.Printf("ListActions err: %v\n", err)
		return "", err
	}
	act, err := result.Recv()
	if err != nil {
		fmt.Printf("recv err: %v\n", err)
		return "", err
	}

	t := act.GetType()
	return t, nil
}
