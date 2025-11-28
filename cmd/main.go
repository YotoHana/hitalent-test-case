package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YotoHana/hitalent-test-case/internal/database"
	"github.com/YotoHana/hitalent-test-case/internal/handlers"
	"github.com/YotoHana/hitalent-test-case/internal/repository"
	"github.com/YotoHana/hitalent-test-case/internal/server"
	"github.com/YotoHana/hitalent-test-case/internal/service"
)

func main() {
	srvCfg := server.DefaultConfig()
	dbCfg := database.DefaultConfig()

	db, err := database.NewDatabase(dbCfg)
	if err != nil {
		log.Fatalf("failed create gorm: %v", err)
	}

	questionRepo := repository.NewQuestionRepository(db.DB)
	answerRepo := repository.NewAnswerRepository(db.DB)

	questionService := service.NewQuestionService(questionRepo, answerRepo)
	answerService := service.NewAnswerService(answerRepo)

	questionHandlers := handlers.NewQuestionHandler(questionService)
	answerHandlers := handlers.NewAnswerHandler(answerService)

	srv := server.NewServer(srvCfg, questionHandlers, answerHandlers)

	srv.ImplementHandlers()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("failed start server: %v", err)
		}
	}()

	<-sigChan

	srv.Stop(context.Background())
	db.Close()
}
