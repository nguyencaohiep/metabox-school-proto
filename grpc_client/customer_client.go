package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/customer"

	"google.golang.org/grpc"
)

var (
	_customerClient        *CustomerClientStruct
	loadCustomerClientOnce sync.Once
)

func ConnectToCustomerServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadCustomerClientOnce.Do(func() {
		_customerClient = new(CustomerClientStruct)
		err = _customerClient.Connect(addr, options...)
	})

	return err
}

func CustomerClient() *CustomerClientStruct {
	if _customerClient == nil {
		panic("grpc customer client: like client is not initiated")
	}

	return _customerClient
}

type CustomerClientStruct struct {
	customer.CustomerServiceClient
	clientConn *grpc.ClientConn
}

func (c *CustomerClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.CustomerServiceClient = customer.NewCustomerServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *CustomerClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
