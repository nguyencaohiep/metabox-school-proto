package grpc_client

import (
	"sync"

	"github.com/nguyencaohiep/metabox-school-proto/golang/school"
	"google.golang.org/grpc"
)

var (
	_schoolClient        *SchoolClientStruct
	loadSchoolClientOnce sync.Once
)

func ConnectToSchoolServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadSchoolClientOnce.Do(func() {
		_schoolClient = new(SchoolClientStruct)
		err = _schoolClient.Connect(addr, options...)
	})

	return err
}

func SchoolClient() *SchoolClientStruct {
	if _schoolClient == nil {
		panic("grpc school client: like client is not initiated")
	}

	return _schoolClient
}

type SchoolClientStruct struct {
	school.SchoolServiceClient
	clientConn *grpc.ClientConn
}

func (c *SchoolClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	schoolConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.SchoolServiceClient = school.NewSchoolServiceClient(schoolConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = schoolConn
	return nil
}

func (c *SchoolClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
