package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Fact0RR/RTULab/config"
	"github.com/Fact0RR/RTULab/internal/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
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
	pathfine := "excel/violations_fine.xlsx"
	pathSitiz:="excel/citizen_data.xlsx"

	_, err := excelize.OpenFile(pathfine)
	if err != nil {
		log.Panic(err)
		return nil
	}
	_, err = excelize.OpenFile(pathSitiz)
	if err != nil {
		log.Panic(err)
		return nil
	}

	s.Store = store.New(s.Conf.DataBaseString,pathfine,pathSitiz)
	
	if s.Store.Open(s.Conf.K,s.Conf.J) != nil {
		log.Fatal("Подключение не открыто ", s.Store.Open(s.Conf.K,s.Conf.J))
	}
	s.ConfigureRouter()
	return http.ListenAndServe(s.Conf.Port, handlers.LoggingHandler(os.Stdout, s.Router))
}
