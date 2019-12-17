package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	bindAddr = flag.String("bind", "localhost:58001", "Bind address to listen for incoming requests")
	status   = flag.Int("status", http.StatusOK, "Change the status code to return")
)

func logToStdoutAndReturn(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(*status)
	buf := bytes.Buffer{}
	buf.WriteString("\n----- GITHUB ACTION REQUEST -----\n")
	req.WriteProxy(&buf)
	buf.WriteString("\n----- END GITHUB ACTION REQUEST -----\n")
	io.Copy(os.Stdout, &buf)
}

func main() {
	flag.Parse()

	log.Printf("Bind addr: %v", *bindAddr)

	if err := http.ListenAndServe(*bindAddr, http.HandlerFunc(logToStdoutAndReturn)); err != nil {
		log.Printf("Opsie.... %v", err)
	}
}
