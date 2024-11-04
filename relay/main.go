package main

// there is not much to modify in this code
// the relay seems to  be ready
import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"

	"github.com/ranon-rat/frensmatria/common"
)

var (
	addresses   = make(map[string]string)
	connections = make(map[string]net.Conn)
)

func RandomConnectionsID() (nodes []string) {
	existent := make(map[string]bool)
	for len(nodes) < min(len(connections), 10) {
		for id := range connections {
			if rand.Float32() < max(0.5/float32(len(connections)), 0.1) {
				if existent[id] {
					continue

				}
				nodes = append(nodes, id)
				existent[id] = true
				for len(nodes) >= min(len(connections), 10) {
					return nodes
				}
			}
		}
	}
	return nodes
}
func manageConnections(c net.Conn) {

	defer c.Close()
	var initialize common.WantConnect
	reader := json.NewDecoder(c)

	if reader.Decode(&initialize) != nil {

		return
	}

	ID := common.HashSHA256(initialize.SDPOffer)
	fmt.Println(ID)
	defer delete(connections, ID)
	defer delete(addresses, ID)

	json.NewEncoder(c).Encode(common.IDResponse{ID: ID, NodesID: RandomConnectionsID()})

	connections[ID] = c
	addresses[ID] = initialize.SDPOffer
	for {
		var conInterest common.WantConnect

		if reader.Decode(&conInterest) != nil {
			break
		}
		//  in this case the client is interested in someone else sdp
		if conInterest.SDPOffer == "" {
			v, e := addresses[conInterest.IDNode]
			if !e {
				json.NewEncoder(c).Encode(common.Initial{})
				continue
			}
			json.NewEncoder(c).Encode(common.Initial{SDP: v, Kind: common.ConnectTo})
			continue
		}

		// in this case this will update your sdp
		if conInterest.IDNode == "" && conInterest.SDPOffer != "" { // this will add more stuff to it
			addresses[ID] = conInterest.SDPOffer
			continue
		}
		// and this just sends the sdp to the other guy
		cconn, e := connections[conInterest.IDNode]
		if !e {
			// empty
			json.NewEncoder(c).Encode(common.Initial{})
			continue
		}
		json.NewEncoder(cconn).Encode(common.Initial{SDP: conInterest.SDPOffer, Kind: common.Confirm})

	}
}
func main() {

	port := flag.String("port", "9090", "its the port for the local server")
	flag.Parse()

	server, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		panic(err)
	}
	defer server.Close()
	fmt.Println("starting server")
	fmt.Printf("you can connect via %s:%s\n", "localhost", *port)

	for {

		conn, err := server.Accept()
		if err != nil {
			continue
		}

		go manageConnections(conn)
		fmt.Println("new connection", conn.RemoteAddr())

	}
}
