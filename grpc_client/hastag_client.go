package grpc_client

import (
	"sync"

	"google.golang.org/grpc"
	"git.ctisoftware.vn/cti-group/proto-lib/golang/hastag"
)

var (
	_hastagClient        *HastagClientStruct
	loadHastagClientOnce sync.Once
)

func ConnectToHastagServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadHastagClientOnce.Do(func() {
		_hastagClient = new(HastagClientStruct)
		err = _hastagClient.Connect(addr, options...)
	})

	return err
}

func HastagClient() *HastagClientStruct {
	if _hastagClient == nil {
		panic("grpc hastag client: like client is not initiated")
	}

	return _hastagClient
}

type HastagClientStruct struct {
	hastag.HastagServiceClient
	clientConn *grpc.ClientConn
}

func (c *HastagClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.HastagServiceClient = hastag.NewHastagServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *HastagClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
