package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/service"

	"google.golang.org/grpc"
)

var (
	_serviceClient        *ServiceClientStruct
	loadServiceClientOnce sync.Once
)

func ConnectToServiceServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadServiceClientOnce.Do(func() {
		_serviceClient = new(ServiceClientStruct)
		err = _serviceClient.Connect(addr, options...)
	})

	return err
}

func ServiceClient() *ServiceClientStruct {
	if _serviceClient == nil {
		panic("grpc service client: like client is not initiated")
	}

	return _serviceClient
}

type ServiceClientStruct struct {
	service.ServiceServiceClient
	clientConn *grpc.ClientConn
}

func (c *ServiceClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	serviceConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.ServiceServiceClient = service.NewServiceServiceClient(serviceConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = serviceConn
	return nil
}

func (c *ServiceClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
