package server

import (
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func Launch(config WebConfig) {
	verbose = config.Verbose

	if verbose {
		log.Println("DEBUG: Verbose logging enabled")
	}

	router := mux.NewRouter()

	if verbose {
		router.Use(logVerboseMiddleware)
	}

	router.Path("/").Methods("GET").HandlerFunc(getUUID)
	router.Path("/stats").Methods("GET").HandlerFunc(getStats)

	tlsConfig.BuildNameToCertificate()

	s := &http.Server{
		Addr:           config.Address,
		Handler:        router,
		ReadTimeout:    config.TimeoutSeconds * time.Second,
		WriteTimeout:   config.TimeoutSeconds * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsConfig,
	}

	log.Printf("INFO: Starting listener at %s", config.Address)

	if verbose {
		log.Printf("DEBUG: TLS cert file: %s", config.TLSCert)
		log.Printf("DEBUG: TLS key file: %s", config.TLSKey)
	}

	err := s.ListenAndServeTLS(config.TLSCert, config.TLSKey)

	if err != nil {
		log.Printf("FATAL: Web server exited: %s", err.Error())
	}

}

func logVerboseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("DEBUG: request:", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func getUUID(w http.ResponseWriter, r *http.Request) {
	// generate UUID
	rawUUID, err := exec.Command("uuidgen").Output()
	uuid := string(rawUUID[:len(rawUUID)-1]) // strip the trailing newline

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("FATAL: Failed to generate UUID")
		panic("Failed to generate UUID")
	}

	if verbose {
		// log UUID
		log.Printf("INFO: Generated: %s", uuid)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, uuid)
}

func getStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "{}")
}