package nfsserver

import (
	"os"
	"os/exec"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jdinabox/go-await"
	"k8s.io/klog/v2"
)

const (
	defaultConfig = "/data/config/nfs/conf.json"
)

var confPath string

func init() {
	confPath = os.Getenv("NFS_CONFIG")
	if confPath == "" {
		confPath = defaultConfig
	}
}

func NfsServer(ai *await.Interrupt) {
	defer ai.Done()
	// Init nfs
	klog.Info("Setting up nfs")
	// Read config
	config, err := ReadConfig(confPath)
	logger.Fatal(err)

	// Generate /etc/exports file from config
	exports := config.GenExports()
	// Write exports to /etc/exports
	os.WriteFile("/etc/exports", []byte(exports), 0644)

	// Run exportfs -arv
	run("exportfs", "-arv")
	// Run rpcbind
	run("rpcbind")
	// Run rpc.statd
	run("rpc.statd")
	// Run rpc.nfsd
	run("rpc.nfsd")
	// Run rpc.mountd
	run("rpc.mountd")

	// Get configured exports
	if klog.V(4).Enabled() {
		o, err := exec.Command("exportfs", "-v").Output()
		logger.Error(err)
		klog.Info("Exports from exportfs -v\n" + string(o))
	}

	// Wait for system interrupt
	ai.Await()
}

func run(name string, args ...string) {
	if klog.V(4).Enabled() {
		klog.Info("Running "+name+" ", args)
	}
	logger.Fatal(exec.Command(name, args...).Run())
}
