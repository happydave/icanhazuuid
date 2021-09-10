package server

import (
	"crypto/tls"
	"time"
)

var verbose bool

type WebConfig struct {
	Verbose        bool
	Address        string
	TLSCert        string
	TLSKey         string
	TimeoutSeconds time.Duration
}

var tlsCiphers = []uint16{
	// TLS 1.3
	tls.TLS_AES_256_GCM_SHA384,
}

var tlsConfig = &tls.Config{
	CipherSuites:             tlsCiphers,
	PreferServerCipherSuites: true,
	SessionTicketsDisabled:   true,
	MinVersion:               tls.VersionTLS13,
	MaxVersion:               tls.VersionTLS13,
}
