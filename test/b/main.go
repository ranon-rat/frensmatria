package main

import (
	"fmt"

	"github.com/pion/webrtc/v3"
)

func main() {
	// Configurar servidor STUN
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	//uh hum
	// Crear la conexión WebRTC
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Crear un canal de datos y manejar su apertura
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Println("Canal de datos recibido en Nodo B")
		d.OnOpen(func() {
			fmt.Println("Canal de datos abierto en Nodo B")
		})

		// Manejar recepción de mensajes
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Mensaje recibido en Nodo B: %s\n", string(msg.Data))
		})
	})

	// Esperar la oferta SDP de Nodo A
	var sdp string
	fmt.Println("Introduce la oferta SDP de Nodo A:")
	fmt.Scanln(&sdp)

	// Establecer la oferta como la descripción remota
	offer := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  sdp,
	}
	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	// Crear una respuesta SDP y establecerla como la descripción local
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	// Enviar la respuesta SDP de vuelta a Nodo A
	fmt.Printf("Respuesta SDP generada:\n%s\n", answer.SDP)

	// Aquí es donde normalmente enviarías la respuesta a Nodo A a través de tu mecanismo de señalización

	select {}
}
