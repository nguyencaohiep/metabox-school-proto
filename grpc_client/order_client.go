package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/order"

	"google.golang.org/grpc"
)

var (
	_orderClient        *OrderClientStruct
	loadOrderClientOnce sync.Once
)

func ConnectToOrderServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadOrderClientOnce.Do(func() {
		_orderClient = new(OrderClientStruct)
		err = _orderClient.Connect(addr, options...)
	})

	return err
}

func OrderClient() *OrderClientStruct {
	if _orderClient == nil {
		panic("grpc OrderClient: client is not initiated")
	}

	return _orderClient
}

type OrderClientStruct struct {
	order.OrderServiceClient
	clientConn *grpc.ClientConn
}

func (c *OrderClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	orderConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.OrderServiceClient = order.NewOrderServiceClient(orderConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = orderConn
	return nil
}

func (c *OrderClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
