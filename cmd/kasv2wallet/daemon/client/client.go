package client

import (
	"context"
	"github.com/kasv2/kasv2d/cmd/kasv2wallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/kasv2/kasv2d/cmd/kasv2wallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kaspawalletd server, and returns the client instance
func Connect(address string) (pb.KaspawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kaspawallet daemon is not running, start it with `kasv2wallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKaspawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
