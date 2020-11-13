package cmd

import (
	"time"

	"github.com/golang/groupcache"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"

	"github.com/dahernan/gopherdiscovery/pkg/discovery"
)

var (
	me              string
	updateUrlServer string
	updateUrlPubSub string
)

var UpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update the peers in groupcache.",
	Long:    "Update the peers in groupcache.",
	Run: func(cmd *cobra.Command, args []string) {

		// opts
		opts := discovery.Options{
			SurveyTime:   1 * time.Second,
			RecvDeadline: 1 * time.Second,
			PollTime:     2 * time.Second,
		}

		// on the server
		server, err := discovery.Server(updateUrlServer, updateUrlPubSub, opts)
		checkErr(err)

		pp.Println("server:", server)

		// any of the peers
		pool := groupcache.NewHTTPPool(me)

		client, err := discovery.ClientWithSub(updateUrlServer, updateUrlPubSub, me)
		checkErr(err)

		peers, err = client.Peers()
		checkErr(err)

		for nodes := range peers {
			pool.Set(nodes...)
		}

	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&updateUrlServer, "url-server", "", "tcp://10.0.0.100:40007", "URL Server address")
	UpdateCmd.Flags().StringVarP(&updateUrlPubSub, "url-pubsub", "", "tcp://10.0.0.100:50007", "URL Pub-SUB service address")
	UpdateCmd.Flags().StringVarP(&me, "me", "", "client0", "Client Name")
	RootCmd.AddCommand(UpdateCmd)
}
