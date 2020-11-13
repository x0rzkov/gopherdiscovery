package cmd

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"

	"github.com/dahernan/gopherdiscovery/pkg/discovery"
)

var (
	peers chan []string
	// nodes             chan []string
	discoverUrlServ   string
	discoverUrlPubSub string
)

var DiscoverCmd = &cobra.Command{
	Use:     "discover",
	Aliases: []string{"d"},
	Short:   "Discover peers in a cluster",
	Long:    "Discover peers in a cluster",
	Run: func(cmd *cobra.Command, args []string) {

		// nodes := make(chan []string)

		opts := discovery.Options{
			SurveyTime:   1 * time.Second,
			RecvDeadline: 1 * time.Second,
			PollTime:     2 * time.Second,
		}

		server, err := discovery.Server(discoverUrlServ, discoverUrlPubSub, opts)
		checkErr(err)
		pp.Println("server:", server)

		// client1
		clientOne, err := discovery.ClientWithSub(discoverUrlServ, discoverUrlPubSub, "client1")
		checkErr(err)

		// client2
		clientTwo, err := discovery.ClientWithSub(discoverUrlServ, discoverUrlPubSub, "client2")
		checkErr(err)

		// client3
		clientThree, err := discovery.ClientWithSub(discoverUrlServ, discoverUrlPubSub, "client3")
		checkErr(err)

		peers, err = clientOne.Peers()
		checkErr(err)

		pp.Println("peers:", peers)
		for nodes := range peers {
			pp.Println("nodes:", nodes)
		}
		// nodes <- peers
		// nodes = []string{"client1", "client2", "client3"}

		// Cancel client2
		clientTwo.Cancel()
		clientThree.Cancel()

		// nodes <- peers
		// nodes = []string{"client1", "client3"}

		// peers, err = clientOne.Peers()
		// for nodes := range peers {
		// 	discovery.AddNodesToCluster(nodes)
		// }

	},
}

func init() {
	DiscoverCmd.Flags().StringVarP(&discoverUrlServ, "url-server", "", "tcp://127.0.0.1:40007", "Server address")
	DiscoverCmd.Flags().StringVarP(&discoverUrlPubSub, "url-pubsub", "", "tcp://127.0.0.1:50007", "Pub-sub address")
	RootCmd.AddCommand(DiscoverCmd)
}
