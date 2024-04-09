package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/role"
	"google.golang.org/grpc"
)

var (
	_roleClient        *RoleClientStruct
	loadRoleClientOnce sync.Once
)

func ConnectToRoleServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadRoleClientOnce.Do(func() {
		_roleClient = new(RoleClientStruct)
		err = _roleClient.Connect(addr, options...)
	})

	return err
}

func RoleClient() *RoleClientStruct {
	if _roleClient == nil {
		panic("grpc role client: like client is not initiated")
	}

	return _roleClient
}

type RoleClientStruct struct {
	role.RoleServiceClient
	clientConn *grpc.ClientConn
}

func (c *RoleClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.RoleServiceClient = role.NewRoleServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *RoleClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
