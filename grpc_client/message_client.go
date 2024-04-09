package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/message"
	"google.golang.org/grpc"
)

var (
	_messageClient        *MessageClientStruct
	loadMessageClientOnce sync.Once
)

func ConnectToMessageServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadMessageClientOnce.Do(func() {
		_messageClient = new(MessageClientStruct)
		err = _messageClient.Connect(addr, options...)
	})

	return err
}

func MessageClient() *MessageClientStruct {
	if _messageClient == nil {
		panic("grpc message client: like client is not initiated")
	}

	return _messageClient
}

type MessageClientStruct struct {
	message.MessageServiceClient
	clientConn *grpc.ClientConn
}

func (c *MessageClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	mesConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.MessageServiceClient = message.NewMessageServiceClient(mesConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = mesConn
	return nil
}

func (c *MessageClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
