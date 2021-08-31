package main

import (
	"flag"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

var (
	configPath    string
	wireguardType string
)

func init() {
	const (
		defaultConfig = "/wireguard/wg0.conf"
		configUsage   = "Path to wg0.conf"
		defaultType   = "s"
		typeUsage     = "Run wireguard as server or client [S]erver/[c]lient"
	)

	flag.StringVar(&configPath, "config", defaultConfig, configUsage)
	flag.StringVar(&configPath, "c", defaultConfig, configUsage+" (shorthand)")

	flag.StringVar(&wireguardType, "type", defaultType, typeUsage)
	flag.StringVar(&wireguardType, "t", defaultType, typeUsage+" (shorthand)")

	flag.Parse()
}

func main() {
	switch strings.ToLower(strings.TrimSpace(wireguardType)) {
	case "s", "server":
		break
	case "c", "client":
		break
	default:
		panic("")
	}

	if err := exec.Command("wg-quick", "up", configPath).Run(); err != nil {
		panic(err.Error())
	}

	// Wait for system interrupt
	serverClose := make(chan struct{})
	go awaitClose(serverClose)
	<-serverClose

	print("\nShutting down")
	print("\n", exec.Command("wg-quick", "down", configPath).Run().Error())
}

func awaitClose(serverClose chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	close(serverClose)
}
