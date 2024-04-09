package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/employee"

	"google.golang.org/grpc"
)

var (
	_employeeClient        *EmployeeClientStruct
	loadEmployeeClientOnce sync.Once
)

func ConnectToEmployeeServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadEmployeeClientOnce.Do(func() {
		_employeeClient = new(EmployeeClientStruct)
		err = _employeeClient.Connect(addr, options...)
	})

	return err
}

func EmployeeClient() *EmployeeClientStruct {
	if _employeeClient == nil {
		panic("grpc employee client: like client is not initiated")
	}

	return _employeeClient
}

type EmployeeClientStruct struct {
	employee.EmployeeServiceClient
	clientConn *grpc.ClientConn
}

func (c *EmployeeClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.EmployeeServiceClient = employee.NewEmployeeServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *EmployeeClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
