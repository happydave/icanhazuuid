package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/happydave/icanhazuuid/server"
)

var version string = "custom"
var build string = "custom"

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	listenIP := flag.String("ip", "", "https listen ip")
	listenPort := flag.Uint("port", 443, "https listen port")
	certOpt := flag.String("cert", "/etc/icanhazuuid/cert.pem", "TLS Certificate")
	keyOpt := flag.String("key", "/etc/icanhazuuid/key.pem", "TLS Key")
	timeoutSecondsOpt := flag.Uint("timeout", 30, "Web server timeout in seconds")
	logVerboseOpt := flag.Bool("v", false, "Verbose logging")
	versionOpt := flag.Bool("version", false, "Show version and exit")
	flag.Parse()

	if *versionOpt {
		log.Printf("Version: %s (%s) built with %s", version, build, runtime.Version())
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *listenIP, *listenPort)

	webConfig := &server.WebConfig{
		Verbose:        *logVerboseOpt,
		Address:        addr,
		TLSCert:        *certOpt,
		TLSKey:         *keyOpt,
		TimeoutSeconds: time.Duration(*timeoutSecondsOpt),
	}

	server.Launch(*webConfig)
}
