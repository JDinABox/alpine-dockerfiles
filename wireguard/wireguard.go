package wireguard

import (
	"log"
	"os"
	"os/exec"
	"strings"

	await "github.com/jdinabox/go-await"
)

const (
	defaultConfig = "/data/wireguard/wg0.conf"
	defaultType   = "s"
)

type WgConf struct {
	ConfPath string
	Type     string
}

func newWgConf() *WgConf {
	w := &WgConf{
		ConfPath: os.Getenv("WG_CONFIG"),
		Type:     os.Getenv("WG_TYPE"),
	}

	w.Defaults()
	w.Normalize()

	return w
}

// Defaults completes empty config vars
func (w *WgConf) Defaults() {
	if w.ConfPath == "" {
		w.ConfPath = defaultConfig
	}
	if w.Type == "" {
		w.Type = defaultType
	}
}

// Normalize config vars
func (w *WgConf) Normalize() {
	w.Type = strings.TrimSpace(w.Type)
	switch strings.ToLower(w.Type) {
	case "server":
		w.Type = "s"
	case "client":
		w.Type = "c"
	}
}

func Wireguard(ai *await.Interrupt) {
	// Add 1 to wait group
	ai.Add(1)
	defer ai.Done()

	w := newWgConf()

	// TODO: Gen wg.conf file
	switch w.Type {
	case "s":
		break
	case "c":
		break
	default:
		panic("Invalid WG_TYPE [[S]erver/[c]lient]")
	}

	// Start wireguard
	if err := exec.Command("wg-quick", "up", w.ConfPath).Run(); err != nil {
		panic(err.Error())
	}

	// Wait for system interrupt
	ai.Await()

	// Stop wireguard
	log.Println("Shutting down")
	log.Println(exec.Command("wg-quick", "down", w.ConfPath).Run().Error())
}
