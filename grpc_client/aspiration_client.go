package grpc_client

import (
	"sync"

	"github.com/nguyencaohiep/metabox-school-proto/golang/aspiration"
	"google.golang.org/grpc"
)

var (
	_aspiratrionClient       *AspirationClientStruct
	loadAspirationClientOnce sync.Once
)

func ConnectToAspirationServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadAspirationClientOnce.Do(func() {
		_aspiratrionClient = new(AspirationClientStruct)
		err = _aspiratrionClient.Connect(addr, options...)
	})

	return err
}

func AspirationClient() *AspirationClientStruct {
	if _aspiratrionClient == nil {
		panic("grpc aspiration client: like client is not initiated")
	}

	return _aspiratrionClient
}

type AspirationClientStruct struct {
	aspiration.AspirationServiceClient
	clientConn *grpc.ClientConn
}

func (c *AspirationClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	aspirationConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.AspirationServiceClient = aspiration.NewAspirationServiceClient(aspirationConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = aspirationConn
	return nil
}

func (c *AspirationClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
