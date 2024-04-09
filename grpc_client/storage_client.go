package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/storage"
	"google.golang.org/grpc"
)

var (
	_storageClient        *StorageClientStruct
	loadStorageClientOnce sync.Once
)

func ConnectToStorageServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadStorageClientOnce.Do(func() {
		_storageClient = new(StorageClientStruct)
		err = _storageClient.Connect(addr, options...)
	})

	return err
}

func StorageClient() *StorageClientStruct {
	if _storageClient == nil {
		panic("grpc storage client: like client is not initiated")
	}

	return _storageClient
}

type StorageClientStruct struct {
	storage.StorageServiceClient
	clientConn *grpc.ClientConn
}

func (c *StorageClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	mesConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.StorageServiceClient = storage.NewStorageServiceClient(mesConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = mesConn
	return nil
}

func (c *StorageClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
