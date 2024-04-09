package grpc_client

import (
	"sync"

	"google.golang.org/grpc"

	rewardPoint "git.ctisoftware.vn/cti-group/proto-lib/golang/reward-point"
)

var (
	_rewardPointClient        *RewardPointClientStruct
	loadRewardPointClientOnce sync.Once
)

func ConnectToRewardPointServer(addr string, options ...grpc.DialOption) error {
	var err error
	loadRewardPointClientOnce.Do(func() {
		_rewardPointClient = new(RewardPointClientStruct)
		err = _rewardPointClient.Connect(addr, options...)
	})

	return err
}

func RewardPointClient() *RewardPointClientStruct {
	if _rewardPointClient == nil {
		panic("grpc rewardPoint client: like client is not initiated")
	}

	return _rewardPointClient
}

type RewardPointClientStruct struct {
	rewardPoint.RewardPointServiceClient
	clientConn *grpc.ClientConn
}

func (c *RewardPointClientStruct) Connect(addr string, options ...grpc.DialOption) error {
	authConn, err := grpc.Dial(addr, options...)
	if err != nil {
		return err
	}

	c.RewardPointServiceClient = rewardPoint.NewRewardPointServiceClient(authConn)
	c.clientConn = new(grpc.ClientConn)
	c.clientConn = authConn
	return nil
}

func (c *RewardPointClientStruct) Close() {
	if c.clientConn == nil {
		return
	}

	defer c.clientConn.Close()
}
