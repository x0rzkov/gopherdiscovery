package cmd

/*
// Using gopherdiscovery to update the peers in groupcache

urlServer := "tcp://10.0.0.100:40007"
urlPubSub := "tcp://10.0.0.100:50007"
me := "http://10.0.0.1"

// on the server
server, err := gopherdiscovery.Server(urlServer, urlPubSub, opts)


// any of the peers
pool := groupcache.NewHTTPPool(me)
client, err := gopherdiscovery.ClientWithSub(urlServer, urlPubSub, me)

peers, err = client.Peers()
for nodes := ranges peers {
	pool.Set(nodes...)
}
*/
