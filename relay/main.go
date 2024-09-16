package main

// there is not much to modify in this code
// the relay seems to  be ready
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

var (
	addresses   = make(map[string]string)
	connections = make(map[string]net.Conn)
)

const (
	Nothing   = 0
	ConnectTo = 1
	Confirm   = 2
)

type Initial struct {
	Kind int    `json:"kind"`
	SDP  string `json:"direction"`
}

// there
type IDResponse struct {
	ID string `json:"id"`
}

type WantConnect struct {
	IDNode   string `json:"idNode"`
	SDPOffer string `json:"SDP"` // if this is empty, that means that i shouldnt send it to the IDNode
}

func HashSHA256(input string) string {
	// Crear un hash SHA-256
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
func manageConnections(c net.Conn) {

	defer c.Close()
	var initialize WantConnect
	reader := json.NewDecoder(c)

	if reader.Decode(&initialize) != nil {

		return
	}

	ID := HashSHA256(initialize.SDPOffer)
	fmt.Println(ID)
	defer delete(connections, ID)
	defer delete(addresses, ID)

	json.NewEncoder(c).Encode(IDResponse{ID: ID})

	connections[ID] = c
	addresses[ID] = initialize.SDPOffer
	for {
		var conInterest WantConnect

		if reader.Decode(&conInterest) != nil {
			break
		}
		//  in this case the client is interested in someone else sdp
		if conInterest.SDPOffer == "" {
			v, e := addresses[conInterest.IDNode]
			if !e {
				json.NewEncoder(c).Encode(Initial{})
				continue
			}
			json.NewEncoder(c).Encode(Initial{SDP: v, Kind: ConnectTo})
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
			json.NewEncoder(c).Encode(Initial{})
			continue
		}
		json.NewEncoder(cconn).Encode(Initial{SDP: conInterest.SDPOffer, Kind: Confirm})

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
	fmt.Printf("you can connect via %s:%s\n", "localhost", port)

	for {

		conn, err := server.Accept()
		if err != nil {
			continue
		}

		go manageConnections(conn)
		fmt.Println("new connection", conn.RemoteAddr())

	}
}
