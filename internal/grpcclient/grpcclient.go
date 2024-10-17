package grpcclient

import (
	proto "github.com/sletkov/thumbnail-proxy-client/pkg/sdk/go/thumbnailproxy_grpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (proto.ThumbnailProxyClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(viper.GetString("grpc-server-address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := proto.NewThumbnailProxyClient(conn)

	return client, conn, nil
}
