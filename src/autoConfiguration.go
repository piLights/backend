package main

import (
	"encoding/json"
	"log"
	"net"
)

type configurationRequest struct {
	DeviceID string
}

type configuration struct {
	IPAddress net.IP
	Port      int
	Version   string
}

//startServer starts the GRPC-server and binds to the defined address
func startAutoConfigurationServer() {
	if *debug {
		log.Printf("Binding to %s\n", *bindTo)
	}

	socket, error := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 13338,
	})
	if error != nil {
		log.Fatalf("Failed to start the background-configuration server on %s, port %d\n", net.IPv4zero.String(), 13338)
	}

	for {
		listen(socket)
	}
}

func listen(socket *net.UDPConn) {
	data := make([]byte, 4096)
	length, remoteAddr, error := socket.ReadFromUDP(data)
	if error != nil {
		log.Fatal(error)
	}

	//@ToDO: Check, if the user wants the file
	var request configurationRequest
	error = json.Unmarshal(data[:length], &request)
	if error != nil {
		log.Fatal(error)
	}
}
