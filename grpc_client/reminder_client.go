package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/reminder"
	"google.golang.org/grpc"
)

var (
	_reminderClient        *ReminderClientStruct
	loadReminderClientOnce sync.Once
)

func ConnectToReminderServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadReminderClientOnce.Do(func() {
		_reminderClient = new(ReminderClientStruct)
		err = _reminderClient.Connect(addr, options...)
	})

	return err
}

func ReminderClient() *ReminderClientStruct {
	if _reminderClient == nil {
		panic("grpc reminder client: like client is not initiated")
	}

	return _reminderClient
}

type ReminderClientStruct struct {
	reminder.ReminderServiceClient
	clientConn *grpc.ClientConn
}

func (c *ReminderClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	mesConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.ReminderServiceClient = reminder.NewReminderServiceClient(mesConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = mesConn
	return nil
}

func (c *ReminderClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
