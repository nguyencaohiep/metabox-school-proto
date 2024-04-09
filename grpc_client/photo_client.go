package grpc_client

import (
	"git.ctisoftware.vn/cti-group/proto-lib/golang/photo"
	"sync"

	"google.golang.org/grpc"
)

var (
	_photoClient        *PhotoClientStruct
	loadPhotoClientOnce sync.Once
)

func ConnectToPhotoServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadPhotoClientOnce.Do(func() {
		_photoClient = new(PhotoClientStruct)
		err = _photoClient.Connect(addr, options...)
	})

	return err
}

func PhotoClient() *PhotoClientStruct {
	if _photoClient == nil {
		panic("grpc PhotoClient: client is not initiated")
	}

	return _photoClient
}

type PhotoClientStruct struct {
	photo.PhotoServiceClient
	clientConn *grpc.ClientConn
}

func (c *PhotoClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	photoConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.PhotoServiceClient = photo.NewPhotoServiceClient(photoConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = photoConn
	return nil
}

func (c *PhotoClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
