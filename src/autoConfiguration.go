package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"unicode"

	"github.com/davecheney/mdns"
)

//removeWhitespaces removes all whitespaces from the given string
func removeWhitespaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

//startServer starts the GRPC-server and binds to the defined address
func startAutoConfigurationServer() {
	if *debug {
		logChan <- fmt.Sprintf("Binding to %s", *bindTo)
	}

	_, port, error := net.SplitHostPort(*bindTo)
	if error != nil {
		log.Fatal(error)
	}

	//Publish the ServerName
	publishRecord(`_dioder._tcp.local. 60 IN TXT "` + *serverName + `"`)

	//Register _dioder._tcp on the local mDNS domain
	publishRecord("_services._dns-sd._udp.local. 60 IN PTR dioder._tcp.local.")

	cleanHostName := removeWhitespaces(*serverName)
	//A record for servername.local for every IPv4 address
	//AAAA record for serverName.local for every IPv6 address
	publishARecords(cleanHostName)

	//@ToDo: PTR for every IP to serverName.local

	// SRV -> _dioder._tcp.local 60 IN SRV 0 0 PORT HOST
	createSRVRecord(cleanHostName, port)
}

//publishARecords publishes an A or AAAA record on the given hostname with every interface-address
func publishARecords(hostName string) {
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

				publishRecord(hostName + ".local. 60 IN A " + ipAddress.String())
			}
		}
	}
}

//createSRVRecord creates an SRV record announcing the service on the given host:port
func createSRVRecord(hostName, port string) {
	var srvRecord = "_dioder._tcp.local. 60 IN SRV 0 0 " + port + " " + hostName + ".local."
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
