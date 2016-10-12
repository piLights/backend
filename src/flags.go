package main

import "github.com/urfave/cli"

var applicationFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "writeConfiguration",
		Usage: "Write the current configuration to the given file",
	},
	cli.StringFlag{
		Name:        "configurationFile, c",
		Usage:       "Path to the configuration-file",
		Destination: &DioderConfiguration.ConfigurationFile,
	},
	cli.BoolFlag{
		Name:  "update, u",
		Usage: "Fetch the newest version",
	},
	cli.StringFlag{
		Name:        "redPin, r",
		Value:       "18",
		Usage:       "Number of the red pin",
		Destination: &DioderConfiguration.Pins.Red,
	},
	cli.StringFlag{
		Name:        "bluePin, b",
		Value:       "18",
		Usage:       "Number of the blue pin",
		Destination: &DioderConfiguration.Pins.Blue,
	},
	cli.StringFlag{
		Name:        "greenPin, g",
		Value:       "18",
		Usage:       "Number of the green pin",
		Destination: &DioderConfiguration.Pins.Green,
	},
	cli.StringFlag{
		Name:        "bindTo, i",
		Value:       ":13337",
		Usage:       "Address and port to listen on, defaults to 0.0.0.0:13337",
		Destination: &DioderConfiguration.BindTo,
	},
	cli.BoolFlag{
		Name:        "debug, d",
		Usage:       "Turn on the debug-mode",
		Destination: &DioderConfiguration.Debug,
	},
	cli.StringFlag{
		Name:        "password, P",
		Usage:       "The password to protect the endpoint",
		Destination: &DioderConfiguration.Password,
	},
	cli.StringFlag{
		Name:        "piBlaster, p",
		Usage:       "Location of the piBlaster FIFO-file",
		Destination: &DioderConfiguration.PiBlaster,
	},
	cli.StringFlag{
		Name:        "serverName, n",
		Usage:       "The name of the server",
		Value:       "Dioder Server",
		Destination: &DioderConfiguration.ServerName,
	},
	cli.StringFlag{
		Name:        "updateURL",
		Usage:       "Fetch the update from the given URL",
		Value:       UPDATEURL,
		Destination: &DioderConfiguration.UpdateURL,
	},
	cli.BoolFlag{
		Name:        "ipv4Only, 4",
		Usage:       "Enables only IPv4",
		Destination: &DioderConfiguration.IPv4Only,
	},
	cli.BoolFlag{
		Name:        "ipv6Only, 6",
		Usage:       "Enables only IPv6",
		Destination: &DioderConfiguration.IPv6Only,
	},
	cli.BoolFlag{
		Name:        "useAvahi",
		Usage:       "Uses Avahi to serve mDNS",
		Destination: &DioderConfiguration.UseAvahi,
	},
}
