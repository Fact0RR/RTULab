package app

import "net/http"

func (s *Server) StopReportringPeriod(w http.ResponseWriter, r *http.Request){
	err:= s.Store.CallStopReportingPeriod()
	if err != nil{
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("New reporting period started"))
}