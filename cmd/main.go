package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YotoHana/hitalent-test-case/internal/database"
	"github.com/YotoHana/hitalent-test-case/internal/server"
)

func main() {
	srvCfg := server.DefaultConfig()
	dbCfg := database.DefaultConfig()

	srv := server.NewServer(srvCfg)
	db, err := database.NewDatabase(dbCfg)
	if err != nil {
		log.Fatalf("failed create gorm: %v", err)
	}
	

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func ()  {
		if err := srv.Start(); err != nil {
			log.Fatalf("failed start server: %v", err)
		}
	}()

	<- sigChan

	srv.Stop(context.Background())
	db.Close()
}