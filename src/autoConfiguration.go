package main

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
)

type configurationRequest struct {
	DeviceID     string
	CallbackPort int
}

type configuration struct {
	IPAddress net.IP
	Port      string
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
		if *debug {
			log.Printf("Failed to start the background-configuration server on %s, port %d\n", net.IPv4zero.String(), 13338)
		}
		log.Fatal(error)
	}

	defer socket.Close()

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

	remoteAddressHostPort := remoteAddr.String() + strconv.Itoa(request.CallbackPort)
	connection, error := net.Dial("tcp", remoteAddressHostPort)
	if error != nil {
		if *debug {
			log.Printf("Could not connect to %s to send the configuration.\n", remoteAddressHostPort)
		}

		log.Println(error)
	}

	defer connection.Close()

	host, port, error := net.SplitHostPort(*bindTo)
	if error != nil {
		if *debug {
			log.Printf("Could not split host and port: %s", *bindTo)
		}
		log.Fatal(error)
	}
	clientConfiguration := configuration{
		net.ParseIP(host),
		port,
		version,
	}

	//Build the request
	message, error := json.Marshal(clientConfiguration)
	if error != nil {
		log.Fatal(error)
	}

	connection.Write([]byte(message))

}
