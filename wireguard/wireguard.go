package wireguard

import (
	"os"
	"os/exec"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/logger"
	await "github.com/jdinabox/go-await"
	"k8s.io/klog/v2"
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
	klog.Info("Starting wireguard")
	output, err := exec.Command("wg-quick", "up", w.ConfPath).CombinedOutput()
	if klog.V(3).Enabled() {
		klog.Info(string(output))
	}
	logger.Fatal(err)

	// Wait for system interrupt
	ai.Await()

	// Stop wireguard
	klog.Info("Stopping wireguard")
	logger.Error(exec.Command("wg-quick", "down", w.ConfPath).Run())
}
