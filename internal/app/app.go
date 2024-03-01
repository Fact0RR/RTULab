package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Fact0RR/RTULab/config"
	"github.com/Fact0RR/RTULab/internal/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Conf     *config.Config
	Router   *mux.Router
	Store    *store.Store
}

func New() *Server {
	return &Server{
		Conf:     config.GetConfig(),
	}
}

func (s *Server) StartApp() error {
	s.Store = store.New(s.Conf.DataBaseString)

	if s.Store.Open() != nil {
		log.Fatal("Подключение не открыто ", s.Store.Open())
	}
	s.ConfigureRouter()
	return http.ListenAndServe(s.Conf.Port, handlers.LoggingHandler(os.Stdout, s.Router))
}
