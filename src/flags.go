package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	dioderAPI = kingpin.New("dioderAPI", "An RPC interface to Ikea-Dioder")

	bindTo            = dioderAPI.Flag("bindTo", "Address and port to listen on, defaults to 0.0.0.0:13337").Default(":13337").Short('p').String()
	redPin            = dioderAPI.Flag("red", "Number of the red-pin").Default("18").OverrideDefaultFromEnvar("dioderRedPin").Short('r').String()
	greenPin          = dioderAPI.Flag("green", "Number of the green-pin").Default("4").OverrideDefaultFromEnvar("dioderGreenPin").Short('g').String()
	bluePin           = dioderAPI.Flag("blue", "Number of the blue-pin").Default("17").OverrideDefaultFromEnvar("dioderBluePin").Short('b').String()
	debug             = dioderAPI.Flag("debug", "Debug-mode").Bool()
	doUpdate          = dioderAPI.Flag("update", "Update the program to the latest version").Bool()
	updateFromURL     = dioderAPI.Flag("updateURL", "The URL to fetch the new version from").Default(UPDATEURL).String()
	configurationFile = dioderAPI.Flag("configurationFile", "The file to configure").ExistingFile()
	password          = dioderAPI.Flag("password", "The password to protect the endpoint").String()
	cpuProfile        = dioderAPI.Flag("cpuProfile", "Write the CPU-Profile to a file").String()
	piBlaster         = dioderAPI.Flag("piBlaster", "Location of the piBlaster FIFO-file").ExistingFile()
)
