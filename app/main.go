package main

import (
	"flag"
	"fmt"
	"log/slog"
	"nerd3/ise_mock/mock"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// ustawienia gdy chcę logować w json
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger) // Ustawienie jako logger domyślny

	slog.Info("no cześć")
	defer slog.Info("papa, biedaczyska")

	// setting config
	defaultServer := "127.0.0.1"
	defaultPort := "9060"

	listenPortArg := flag.String("p", defaultPort, "listen port")
	listenServer := flag.String("s", defaultServer, "listen server")
	flag.Parse()

	if serverEnv := os.Getenv("SERVER_ADDRESS"); serverEnv != "" {
		*listenServer = serverEnv
	}
	if portEnv := os.Getenv("SERVER_PORT"); portEnv != "" {
		*listenPortArg = portEnv
	}
	listenPort, err := strconv.Atoi(*listenPortArg)
	if err != nil {
		msg := fmt.Sprintf("Server port must be a number. Arg value='%s'", *listenPortArg)
		slog.Error(msg)
		os.Exit(1)
	}
	// end of setting config

	server := http.NewServeMux()

	iseMock := mock.NewMockMux()
	server.Handle("/ers/config/", http.StripPrefix("/ers/config", iseMock))
	server_string := fmt.Sprintf("%s:%d", *listenServer, listenPort)
	slog.Info("Server listen on " + server_string)
	err = http.ListenAndServe(server_string, server)
	if err != nil {
		slog.Error("Error occurred", "message", err)
		os.Exit(1)
	}
}
