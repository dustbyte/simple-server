package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"io"

	log "github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()

	log.Info("Received new request from ", req.UserAgent())

	if err != nil {
		log.Error("Couldn't get hostname")
		w.WriteHeader(500)
		return
	}

	fmt.Fprintln(w, hostname)
}

func main() {
	port := flag.String("port", "80", "Listen port")
	debug := flag.Bool("debug", false, "Set the log level to debug")
	logFile := flag.String("log-file", "", "Log File to send to logs to on top of stdout")

	flag.Parse()

	if *logFile != "" {
		outputFile, err := os.Create(*logFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer outputFile.Close()
		output := io.MultiWriter(os.Stdout, outputFile)
		log.SetOutput(output)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	http.HandleFunc("/", hello)

	log.Info("Running on port ", *port)
	http.ListenAndServe(":" + *port, nil)
}
