package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) ConfigureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	//employees
	router.HandleFunc("/login", s.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/registration", s.RegistrationHandler).Methods(http.MethodPost)
	router.HandleFunc("/work", s.LoginMiddleWare(s.VerifiedMiddleWare(s.GetExcessesFromPool))).Methods(http.MethodGet)
	router.HandleFunc("/work", s.LoginMiddleWare(s.VerifiedMiddleWare(s.SetAnswerOfExcess))).Methods(http.MethodPost)
	router.HandleFunc("/lern",s.LoginMiddleWare(s.VerifiedMiddleWare(s.GetLern))).Methods(http.MethodGet)
	//router.HandleFunc("/test", s.LoginMiddleWare(s.VerifiedMiddleWare(testHandler))).Methods(http.MethodGet)
	//cameras
	router.HandleFunc("/reg_cam",s.RegCamHandler).Methods(http.MethodPost)
	router.HandleFunc("/send", s.SendExcess).Methods(http.MethodPost)
	//admin
	router.HandleFunc("/loginA",s.LoginAdminHandler).Methods(http.MethodPost)
	router.HandleFunc("/verify", s.LoginAdminMiddleWare(s.GetNoVerify)).Methods(http.MethodGet)
	router.HandleFunc("/verify", s.LoginAdminMiddleWare(s.VerifyEmpoloyee)).Methods(http.MethodPost)
	router.HandleFunc("/statistic", s.GetCurrentStatistic).Methods(http.MethodGet)
	router.HandleFunc("/analitic", s.GetAnalitic).Methods(http.MethodPost)
	router.HandleFunc("/end_reporting_period",s.LoginAdminMiddleWare(s.StopReportringPeriod)).Methods(http.MethodGet)
	//router.HandleFunc("/tA",s.LoginAdminMiddleWare(testAdminHandler))

	s.Router = router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello guest!!"))
}

func testHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user!!"))
}

func testAdminHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin!!"))
}