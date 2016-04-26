package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	bindTo        = kingpin.Flag("bindTo", "Address and port to listen on, defaults to 0.0.0.0:13337").Default(":13337").Short('p').String()
	redPin        = kingpin.Flag("red", "Number of the red-pin").Default("18").OverrideDefaultFromEnvar("dioderRedPin").Short('r').Int()
	greenPin      = kingpin.Flag("green", "Number of the green-pin").Default("4").OverrideDefaultFromEnvar("dioderGreenPin").Short('g').Int()
	bluePin       = kingpin.Flag("blue", "Number of the blue-pin").Default("17").OverrideDefaultFromEnvar("dioderBluePin").Short('b').Int()
	debug         = kingpin.Flag("debug", "Debug-mode").Bool()
	doUpdate      = kingpin.Flag("update", "Update the program to the latest version").Bool()
	updateFromURL = kingpin.Flag("updateURL", "The URL to fetch the new version from").String()
)
