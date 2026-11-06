package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	srv := server.New(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
