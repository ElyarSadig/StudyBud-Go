package main

import (
	"net/http"

	"github.com/elyarsadig/studybud-go/transport"
)

func main() {
	server := transport.NewHTTPServer(":8080")
	server.AddHandler(transport.GET, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Everything is Ok!"))
	})
	server.Start()
	<-server.Notify()
}
