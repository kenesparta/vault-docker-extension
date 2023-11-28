package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	var socketPath string
	flag.StringVar(
		&socketPath,
		"socket",
		"/run/guest/volumes-service.sock",
		"Unix domain socket to listen on",
	)
	flag.Parse()

	err := os.RemoveAll(socketPath)
	if err != nil {
		log.Fatal(err)
	}

	logrus.New().Infof("Starting listening on %s\n", socketPath)
	router := echo.New()
	router.HideBanner = true
	startURL := ""

	ln, err := listen(socketPath)
	if err != nil {
		log.Fatal(err)
	}
	router.Listener = ln
	router.POST("/vault", vault)
	log.Fatal(router.Start(startURL))
}

func listen(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}
