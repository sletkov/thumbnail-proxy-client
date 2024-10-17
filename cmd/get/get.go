package get

import (
	"context"
	"fmt"
	"github.com/sletkov/thumbnail-proxy-client/internal/grpcclient"
	proto "github.com/sletkov/thumbnail-proxy-client/pkg/sdk/go/thumbnailproxy_grpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sync"
)

var async bool

func init() {
	GetCmd.PersistentFlags().BoolVarP(&async, "async", "a", false, "verbose output")
	viper.BindPFlag("async", GetCmd.PersistentFlags().Lookup("async"))
}

var GetCmd = &cobra.Command{
	Use:   "get [youtube urls...]",
	Short: "Get youtube video thumbnail by link ",
	Long:  `All software has versions. This is Hugo's`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, conn, err := grpcclient.New()
		defer conn.Close()

		if err != nil {
			fmt.Printf("failed to create GRPC client: %v", err)
			return
		}

		resp, err := client.GetThumbnail(context.Background(), &proto.URLRequest{URL: args})
		if err != nil {
			fmt.Println(err)
		}

		urls := make([]string, 0, len(resp.URL))
		for _, url := range resp.URL {
			urls = append(urls, url)
		}

		if async {
			var wg sync.WaitGroup
			for _, url := range urls {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					defer fmt.Println("File was successfully downloaded!")
					downloadFile(url)
				}(url)

			}
			wg.Wait()
		} else {
			for _, url := range urls {
				downloadFile(url)
				fmt.Println("File was successfully downloaded!")
			}
		}
	},
}
