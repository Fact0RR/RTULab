package app

import (
	"encoding/json"
	"net/http"

	"github.com/Fact0RR/RTULab/internal/model/analitic"
)

type AnalitiRequest struct{
	Id int `json:"id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}

func (s *Server) GetAnalitic(w http.ResponseWriter, r *http.Request) {
	var ar AnalitiRequest
	json.NewDecoder(r.Body).Decode(&ar)

	dirtyData,err:= s.Store.GetAnaliticFromDB(ar.Id,ar.StartDate,ar.EndDate)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	coutUnknown,err:= s.Store.GetCountUnkownFromDB(ar.Id)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	cA:=analitic.GetCleanAnalitic(dirtyData,coutUnknown)

	b, err := json.Marshal(cA)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
	w.Write(b)

}