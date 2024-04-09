package grpc_client

import (
	"sync"

	"git.ctisoftware.vn/cti-group/proto-lib/golang/search"
	"google.golang.org/grpc"
)

var (
	_searchClient        *SearchClientStruct
	loadSearchClientOnce sync.Once
)

func ConnectToSearchServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadSearchClientOnce.Do(func() {
		_searchClient = new(SearchClientStruct)
		err = _searchClient.Connect(addr, options...)
	})

	return err
}

func SearchClient() *SearchClientStruct {
	if _searchClient == nil {
		panic("grpc search client: like client is not initiated")
	}

	return _searchClient
}

type SearchClientStruct struct {
	search.SearchServiceClient
	clientConn *grpc.ClientConn
}

func (c *SearchClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.SearchServiceClient = search.NewSearchServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *SearchClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
