package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) ConfigureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", s.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/registration", s.RegistrationHandler).Methods(http.MethodPost)
	router.HandleFunc("/test", s.LoginMiddleWare(testHandler)).Methods(http.MethodGet)

	s.Router = router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello guest!!"))
}

func testHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user!!"))
}