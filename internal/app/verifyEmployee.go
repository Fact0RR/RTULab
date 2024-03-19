package app

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type DataVerify struct{
	Login string `json:"login"`
	Verify bool `json:"verify"`
}

func (s *Server) VerifyEmpoloyee(w http.ResponseWriter, r *http.Request){
	var dv DataVerify
	json.NewDecoder(r.Body).Decode(&dv)
	id,err := s.Store.SetVerify(dv.Login,dv.Verify)
	if err !=nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(strconv.Itoa(id)))
}

func (s *Server) GetNoVerify(w http.ResponseWriter, r *http.Request){
	nvu,err := s.Store.GetNonVerified()
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