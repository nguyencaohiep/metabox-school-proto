package grpc_client

import (
	"sync"

	"github.com/nguyencaohiep/metabox-school-proto/golang/authenticator"
	"google.golang.org/grpc"
)

var (
	_authenticatorClient        *AuthenticatorClientStruct
	loadAuthenticatorClientOnce sync.Once
)

func ConnectToAuthenticatorServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadAuthenticatorClientOnce.Do(func() {
		_authenticatorClient = new(AuthenticatorClientStruct)
		err = _authenticatorClient.Connect(addr, options...)
	})

	return err
}

func AuthenticatorClient() *AuthenticatorClientStruct {
	if _authenticatorClient == nil {
		panic("grpc authenticator client: like client is not initiated")
	}

	return _authenticatorClient
}

type AuthenticatorClientStruct struct {
	authenticator.AuthenticatorServiceClient
	clientConn *grpc.ClientConn
}

func (c *AuthenticatorClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.AuthenticatorServiceClient = authenticator.NewAuthenticatorServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *AuthenticatorClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
