package main

import (
	"fmt"
	"log"
	"net"

	"github.com/davecheney/mdns"
)

//startServer starts the GRPC-server and binds to the defined address
func startAutoConfigurationServer() {
	if *debug {
		logChan <- fmt.Sprintf("Binding to %s", *bindTo)
	}

	publishRecord(`_dioder._tcp.local. 60 IN TXT "` + *serverName + `"`)

	host, port, error := net.SplitHostPort(*bindTo)
	if error != nil {
		log.Fatal(error)
	}

	if host != "" {
		createSRVRecord(host, port)
	} else {
		//If no host is found, create an SRV-record for every IP on this machine
		addressList, error := net.InterfaceAddrs()
		if error != nil {
			log.Fatal(error)
		}

		for _, address := range addressList {
			ipnet, ok := address.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.String() != "" {
					// Do not publish IPv4 records if IPv4 is disabled
					if *ipv4only && ipnet.IP.To4() == nil {
						continue
					}

					// Do not publish IPv6 records if IPv6 is disabled
					if *ipv6only && ipnet.IP.To4() != nil {
						continue
					}

					ipAddress, _, error := net.ParseCIDR(address.String())
					if error != nil {
						log.Fatal(error)
					}

					createSRVRecord(ipAddress.String(), port)
				}
			}
		}
	}
}

//createSRVRecord creates an SRV record announcing the service on the given host:port
func createSRVRecord(host, port string) {
	var srvRecord = "_dioder._tcp.local. 60 IN SRV 0 0 " + port + " " + host
	publishRecord(srvRecord)
}

//publishRecord publishes an record
func publishRecord(resourceRecord string) {
	if *debug {
		logChan <- fmt.Sprintf("Setting resourceRecord: %s", resourceRecord)
	}

	error := mdns.Publish(resourceRecord)
	if error != nil {
		log.Fatalf(`Unable to publish record "%s": %v`, resourceRecord, error)
	}
}
