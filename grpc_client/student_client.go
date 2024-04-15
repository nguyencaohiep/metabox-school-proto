package grpc_client

import (
	"sync"

	"github.com/nguyencaohiep/metabox-school-proto/golang/student"
	"google.golang.org/grpc"
)

var (
	_studentClient        *StudentClientStruct
	loadStudentClientOnce sync.Once
)

func ConnectToStudentServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadStudentClientOnce.Do(func() {
		_studentClient = new(StudentClientStruct)
		err = _studentClient.Connect(addr, options...)
	})

	return err
}

func StudentClient() *StudentClientStruct {
	if _studentClient == nil {
		panic("grpc student client: like client is not initiated")
	}

	return _studentClient
}

type StudentClientStruct struct {
	student.StudentServiceClient
	clientConn *grpc.ClientConn
}

func (c *StudentClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	schoolConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.StudentServiceClient = student.NewStudentServiceClient(schoolConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = schoolConn
	return nil
}

func (c *StudentClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
