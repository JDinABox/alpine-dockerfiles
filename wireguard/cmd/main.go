package main

import (
	"github.com/jdinabox/alpine-dockerfiles/wireguard"
	await "github.com/jdinabox/go-await"
)

func main() {
	ai := await.NewInterrupt()

	go wireguard.Wireguard(ai)

	ai.Wait()
}
