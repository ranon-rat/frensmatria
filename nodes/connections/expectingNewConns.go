package connections

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

/*
this function works for mantaining multiple client connections between nodes
its quite useful specially for mantaining some kind of stability

use it in a goroutine
*/
func StartPeer() string {
	// esto va a manejar multiples conexiones
	for {
		// so we create a new channel of data
		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		dataChannel, err := peerConn.CreateDataChannel("data", nil)
		if err != nil {
			panic(err)
		}

		// Manejar la apertura del canal de datos
		// puedo simplemente mantener el dataChannel fuera de esto creo
		dataChannel.OnOpen(func() {
			fmt.Println("Canal de datos abierto en Main-node")
			for {

				err := dataChannel.SendText("sup dude")
				if err != nil {
					break
				}
				time.Sleep(time.Second)
			}
		})

		// Manejar mensajes recibidos en el canal de datos
		dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Mensaje recibido: %s\n", string(msg.Data))
		})

		// Crear la oferta SDP para enviarla a otros nodos
		offer, err := peerConn.CreateOffer(nil)
		if err != nil {
			panic(err)
		}
		gatherComplete := webrtc.GatheringCompletePromise(peerConn)

		// Establecer la oferta como descripción local
		err = peerConn.SetLocalDescription(offer)
		if err != nil {
			panic(err)
		}

		// Esperar a que la recolección de candidatos ICE termine
		<-gatherComplete
		// Mostrar la oferta para que los otros nodos la utilicen
		channels.SDPChanInivitation <- peerConn.LocalDescription().SDP
		// espero a que se me invite a una respuesta
		answerSDP := <-SDPAnswerChan

		answer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeAnswer,
			SDP:  answerSDP,
		}

		if peerConn.SetRemoteDescription(answer) != nil {
			continue
		}

	}
}
