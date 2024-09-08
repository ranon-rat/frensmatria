package main

// there is not much to modify in this code
// the relay seems to  be ready
import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/ranon-rat/frensmatria/core"
)

var addresses = make(map[string]string)
var connections = make(map[string]net.Conn)

func manageConnections(c net.Conn) {

	defer c.Close()
	var initialize core.WantConnect
	reader := json.NewDecoder(c)

	if reader.Decode(&initialize) != nil {

		return
	}

	ID := core.HashSHA256(initialize.SDPOffer)
	fmt.Println(ID)
	defer delete(connections, ID)
	defer delete(addresses, ID)

	json.NewEncoder(c).Encode(core.IDResponse{ID: ID})

	connections[ID] = c
	addresses[ID] = initialize.SDPOffer
	for {
		var conInterest core.WantConnect

		if reader.Decode(&conInterest) != nil {
			break
		}
		// this looks good enough tbh, later i will be making something so it requires a password to enter
		//shushing
		// it should be expecting a new connection so what this will be doing is just checking which ip is the user interested in
		if conInterest.SDPOffer == "" {
			v, e := addresses[conInterest.IDNode]
			if !e {
				json.NewEncoder(c).Encode(core.Initial{})
				continue
			}
			json.NewEncoder(c).Encode(core.Initial{SDP: v, Kind: core.ConnectTo})
			continue
		}
		// deberia de agregar algo para poder conectarme a varios creo

		if conInterest.IDNode == "" && conInterest.SDPOffer != "" { // this will add more stuff to it
			addresses[ID] = conInterest.SDPOffer
			continue
		}
		cconn, e := connections[conInterest.IDNode]
		if !e {
			// empty
			json.NewEncoder(c).Encode(core.Initial{})
			continue
		}
		json.NewEncoder(cconn).Encode(core.Initial{SDP: conInterest.SDPOffer, Kind: core.Confirm})

	}
}
func main() {
	port, e := os.LookupEnv("PORT")
	if !e {
		port = "8080"
	}

	server, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(err)
	}
	defer server.Close()
	fmt.Println("starting server")
	ip, _ := core.GetLocalIP()
	// esto servira sobre todo para poder sacar cierta informacion
	fmt.Printf("you can connect via %s:%s\n", ip, port)

	// ahora envio esto al servidor

	for {

		conn, err := server.Accept()
		if err != nil {
			continue
		}

		go manageConnections(conn)
		fmt.Println("new connection", conn.RemoteAddr())

	}
}
