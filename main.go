package main

import (
	"fmt"
	"github.com/moul/http2curl"
	"log"
	"net/http"
	"os"
)

var port = ":8000"

type GenericHandler struct {
}

func (g GenericHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	request.URL.Scheme = "https"

	fixedPort := port

	if fixedPort[0] == ':' {
		fixedPort = "127.0.0.1" + port
	}

	request.URL.Host = fixedPort

	command, _ := http2curl.GetCurlCommand(request)

	fmt.Println(command)

	if request.URL.Path == "/" || request.URL.Path == "/health" {
		writer.WriteHeader(200)

		return
	}

	writer.WriteHeader(500)
}

func main() {
	if newPort, hasNewPort := os.LookupEnv("LISTEN_ADDR"); hasNewPort {
		port = newPort
	}

	log.Fatal(http.ListenAndServe(port, &GenericHandler{}))
}
