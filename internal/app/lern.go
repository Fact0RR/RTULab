package app

import (
	"encoding/json"
	"net/http"
	
	"strconv"
)


func (s *Server) GetLern(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("p")
	pInt,err := strconv.Atoi(p)
	if err != nil {
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	e := r.URL.Query().Get("e")
	eInt,err := strconv.Atoi(e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	skill := r.URL.Query().Get("s")
	sInt,err := strconv.Atoi(skill)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := r.URL.Query().Get("t")
	ts:= r.URL.Query().Get("ts")
	te:= r.URL.Query().Get("te")
	data,err:= s.Store.GetLernFromDB(pInt,eInt,sInt,t,ts,te)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b,err:= json.Marshal(data)

	w.Write(b)

}
