package core

import (
	"context"
	"fmt"
	"net/http"

	// grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	logging "github.com/ipfs/go-log/v2"
	connmgr "github.com/libp2p/go-libp2p-core/connmgr"
	ma "github.com/multiformats/go-multiaddr"
	threadsClient "github.com/textileio/go-threads/api/client"
	tc "github.com/textileio/go-threads/common"
	threadsNetclient "github.com/textileio/go-threads/net/api/client"
	tutil "github.com/textileio/go-threads/util"
	"google.golang.org/grpc"
)

var (
	log = logging.Logger("core")
)

type Doru struct {
	threadsNetBootStrapper tc.NetBoostrapper

	threadsClient *threadsClient.Client
	threadsNet    *threadsNetclient.Client

	server *grpc.Server
	proxy  *http.Server

	config Config
}

type Config struct {
	Debug bool

	AddressApi         ma.Multiaddr
	AddressThreadsHost ma.Multiaddr
	AddressIpfsHost ma.Multiaddr

	ThreadsConnectionManager connmgr.ConnManager
}

func NewDoru(
	ctx context.Context,
	config Config,
	opts ...Option,
) (*Doru, error) {
	var args Options
	var err error
	for _, opt := range opts {
		opt(&args)
	}

	d := &Doru{
		config: config,
	}

	// Configure textile/threads
	netOptions := []tc.NetOption{
		tc.WithNetHostAddr(config.AddressThreadsHost),
		tc.WithNetDebug(config.Debug),
	}
	// add Badger repo
	// TODO: if this is the only option for persistance currently,
	//       then lets put it in config, instead being of optional
	netOptions = append(netOptions,
		tc.WithNetBadgerPersistence(args.ThreadsBadgerRepoPath))
	if config.ThreadsConnectionManager != nil {
		netOptions = append(netOptions,
		tc.WithConnectionManager(config.ThreadsConnectionManager))
	}
	d.threadsNetBootStrapper, err = tc.DefaultNetwork(netOptions...)
	if err != nil {
		return nil, err
	}

	log.Info("started doru core")
	fmt.Print("hola world")

	return d, nil
}

func (d *Doru) Bootstrap() {
	// TODO: for first scratch code bootstrap to the default textile nodes
	//       but obviously improve this once the code is a bit better formed
	d.threadsNetBootStrapper.Bootstrap(tutil.DefaultBoostrapPeers())
}
