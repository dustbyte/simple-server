package main

import (
	"fmt"
	"os"
	"net/http"
	"flag"
)

var port = flag.String("port", "80", "Listen port")

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Fprintln(w, "HAPPY")
	}

	fmt.Fprintln(w, hostname)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", hello)

	fmt.Println("Running on port", *port)
	http.ListenAndServe(":" + *port, nil)
}
