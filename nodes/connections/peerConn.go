package connections

import (
	"fmt"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/channels"
)

func ONAnswer() {
	for {
		peerConn, err := webrtc.NewPeerConnection(Config)
		if err != nil {
			panic(err)
		}

		// Crear un canal de datos y manejar su apertura
		peerConn.OnDataChannel(func(d *webrtc.DataChannel) {
			fmt.Println("Canal de datos recibido en Nodo B")
			d.OnOpen(func() {
				fmt.Println("Canal de datos abierto en Nodo B")
			})

			// Manejar recepción de mensajes
			d.OnMessage(func(msg webrtc.DataChannelMessage) {
				fmt.Printf("Mensaje recibido en Nodo B: %s\n", string(msg.Data))
			})
		})
		SDP := <-SDPOfferChan
		offer := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  SDP,
		}
		err = peerConn.SetRemoteDescription(offer)
		if err != nil {
			fmt.Println("100", err)

			continue
		}

		// Crear una respuesta SDP
		answer, err := peerConn.CreateAnswer(nil)
		if err != nil {
			fmt.Println("108", err)
			continue

		}
		gatherComplete := webrtc.GatheringCompletePromise(peerConn)
		// Establecer la respuesta como descripción local
		err = peerConn.SetLocalDescription(answer)
		if err != nil {
			fmt.Println("116", err)
			continue
		}

		// Esperar hasta que la recolección de ICE esté completa
		<-gatherComplete
		fmt.Println("everything has passed")
		channels.SDPChanAnswer <- peerConn.LocalDescription().SDP
	}
}
