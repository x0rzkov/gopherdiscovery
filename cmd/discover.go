package cmd

/*
var peers chan []string
urlServer := "tcp://127.0.0.1:40007"
urlPubSub := "tcp://127.0.0.1:50007"

opts := Options{
		SurveyTime:   1 * time.Second,
		RecvDeadline: 1 * time.Second,
		PollTime:     2 * time.Second,
}

server, err := gopherdiscovery.Server(urlServer, urlPubSub, opts)

// client1
clientOne, err := gopherdiscovery.ClientWithSub(urlServer, urlPubSub, "client1")

// client2
clientTwo, err := gopherdiscovery.ClientWithSub(urlServer, urlPubSub, "client2")

// client3
clientThree, err := gopherdiscovery.ClientWithSub(urlServer, urlPubSub, "client3")

peers, err = clientOne.Peers()
nodes <- peers
// nodes = []string{"client1", "client2", "client3"}

// Cancel client2
clientTwo.Cancel()

nodes <- peers
// nodes = []string{"client1", "client3"}

peers, err = clientOne.Peers()
for nodes := range peers {
	AddNodesToCluster(nodes)
}
*/
