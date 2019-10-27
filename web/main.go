package main

// 12:48

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr: *addr,
	}

	fmt.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
