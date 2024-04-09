package grpc_client

import (
	"sync"

	log "git.ctisoftware.vn/cti-group/proto-lib/golang/activity-log"
	"google.golang.org/grpc"
)

var (
	_activityLogClient        *ActivityLogClientStruct
	loadActivityLogClientOnce sync.Once
)

func ConnectToActivityLogServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadActivityLogClientOnce.Do(func() {
		_activityLogClient = new(ActivityLogClientStruct)
		err = _activityLogClient.Connect(addr, options...)
	})

	return err
}

func ActivityLogClient() *ActivityLogClientStruct {
	if _activityLogClient == nil {
		panic("grpc activity log client: like client is not initiated")
	}

	return _activityLogClient
}

type ActivityLogClientStruct struct {
	log.ActivityLogServiceClient
	clientConn *grpc.ClientConn
}

func (c *ActivityLogClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	logConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.ActivityLogServiceClient = log.NewActivityLogServiceClient(logConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = logConn
	return nil
}

func (c *ActivityLogClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
