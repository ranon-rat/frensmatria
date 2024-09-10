package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

func HashSHA256(input string) string {
	// Crear un hash SHA-256
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetLocalIP() (string, error) {
	// Obtener todas las interfaces de red
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Ignorar interfaces down o loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Obtener las direcciones de la interfaz
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			// Verificar si la dirección es una dirección IPv4
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() && !ipnet.IP.IsPrivate() {
				ip := ipnet.IP.String()
				if strings.Contains(ip, ":") {
					return "[" + ip + "]", nil
				}
				return ip, nil
			}
		}
	}

	return "", fmt.Errorf("no se pudo encontrar una dirección IP local")
}
