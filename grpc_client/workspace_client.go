package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/workspace"
	"google.golang.org/grpc"
)

var (
	_workspaceClient        *WorkspaceClientStruct
	loadWorkspaceClientOnce sync.Once
)

func ConnectToWorkspaceServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadWorkspaceClientOnce.Do(func() {
		_workspaceClient = new(WorkspaceClientStruct)
		err = _workspaceClient.Connect(addr, options...)
	})

	return err
}

func WorkspaceClient() *WorkspaceClientStruct {
	if _workspaceClient == nil {
		panic("grpc workspace client: like client is not initiated")
	}

	return _workspaceClient
}

type WorkspaceClientStruct struct {
	workspace.WorkspaceServiceClient
	clientConn *grpc.ClientConn
}

func (c *WorkspaceClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.WorkspaceServiceClient = workspace.NewWorkspaceServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *WorkspaceClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
