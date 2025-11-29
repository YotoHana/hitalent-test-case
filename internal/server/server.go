package server

import (
	"context"
	"net/http"

	"github.com/YotoHana/hitalent-test-case/internal/handlers"
)

type Server struct {
	httpServer       *http.Server
	mux              *http.ServeMux
	questionHandlers *handlers.QuestionHandler
	answerHandlers   *handlers.AnswerHandler
}

func (s *Server) Start() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func (s *Server) ImplementHandlers() {
	s.mux.HandleFunc("/questions", s.questionHandlers.Questions)
	s.mux.HandleFunc("/questions/{id}", s.questionHandlers.QuestionsID)
	s.mux.HandleFunc("/questions/{id}/answers", s.answerHandlers.QuestionsIDAnswers)
	s.mux.HandleFunc("/answers/{id}", s.answerHandlers.AnswersID)
}

func NewServer(
	cfg *Config,
	questionHandlers *handlers.QuestionHandler,
	answerHandlers *handlers.AnswerHandler,
) *Server {
	mux := http.NewServeMux()

	return &Server{
		mux: mux,
		httpServer: &http.Server{
			Addr:         cfg.host + cfg.port,
			Handler:      mux,
			ReadTimeout:  cfg.readTimeout,
			WriteTimeout: cfg.writeTimeout,
			IdleTimeout:  cfg.idleTimeout,
		},
		questionHandlers: questionHandlers,
		answerHandlers:   answerHandlers,
	}
}
