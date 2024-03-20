package app

import (
	"encoding/json"
	"net/http"
)


func (s *Server) GetCurrentStatistic(w http.ResponseWriter, r *http.Request) {
	nvu,err := s.Store.GetCurrentStatisticfromBD()
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(nvu)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
	w.Write(b)
}
