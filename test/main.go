package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/pion/webrtc/v3"
)

// okay creo que tengo una idea de que puedo hacer con esta
// informacion, vamos a intentar implementar esto en nuestro nodo
// vamos a ver si podemos je

func main() {
	// Configuración de la conexión WebRTC con un servidor STUN
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // Servidor STUN
			},
		},
	}

	// Crear la conexión WebRTC
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Crear un canal de datos para intercambiar información
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}

	// Manejar la apertura del canal de datos
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
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Establecer la oferta como la descripción local
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	// Mostrar la oferta para que los otros nodos la utilicen
	fmt.Printf("Oferta SDP generada en Main-node:\n%s\n", offer.SDP)

	// Aquí esperas la respuesta SDP de un nodo secundario
	// Normalmente se usaría un mecanismo de señalización para enviar la oferta y recibir la respuesta

	// Leer la respuesta SDP del nodo secundario
	fmt.Println("Introduce la respuesta SDP del otro nodo:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answerSDP := scanner.Text()

	// Establecer la respuesta como la descripción remota
	answer := webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  answerSDP,
	}
	err = peerConnection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}

	// Mantener la conexión abierta
	select {}
}
