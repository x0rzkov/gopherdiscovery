package cmd

import (
	"context"
	"time"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"

	"github.com/dahernan/gopherdiscovery/pkg/discovery"
)

var (
	clients           []string
	suscribeUrlServ   string
	suscribeUrlPubSub string
)

var SuscribeCmd = &cobra.Command{
	Use:     "suscribe",
	Aliases: []string{"s"},
	Short:   "Subscribe to clients changes.",
	Long:    "Subscribe to clients changes (new connections/disconnections).",
	Run: func(cmd *cobra.Command, args []string) {

		defaultOpts := discovery.Options{
			SurveyTime:   1 * time.Second,
			RecvDeadline: 1 * time.Second,
			PollTime:     2 * time.Second,
		}

		server, err := discovery.Server(suscribeUrlServ, suscribeUrlPubSub, defaultOpts)
		checkErr(err)

		pp.Println("server:", server)

		ctx, cancel := context.WithCancel(context.Background())
		sub, err := discovery.NewSubscriber(ctx, suscribeUrlPubSub)
		checkErr(err)

		discovery.Client(suscribeUrlServ, "client1")
		discovery.Client(suscribeUrlServ, "client2")

		clients = <-sub.Changes()
		// clients = []string{"client1", "client2"}

		discovery.Client(suscribeUrlServ, "client3")

		clients = <-sub.Changes()
		// clients = []string{"client1", "client2", "client3"}

		cancel() // stops subscribe

	},
}

func init() {
	SuscribeCmd.Flags().StringVarP(&suscribeUrlServ, "url-server", "", "tcp://127.0.0.1:40009", "URL server address")
	SuscribeCmd.Flags().StringVarP(&suscribeUrlPubSub, "url-pubsub", "", "tcp://127.0.0.1:50009", "URL Pub-Sub service address")
	RootCmd.AddCommand(SuscribeCmd)
}
