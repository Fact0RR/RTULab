package app

import (
	"encoding/json"
	"net/http"

	"github.com/Fact0RR/RTULab/internal"
)

func (s*Server) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reg internal.RegStruct
	json.NewDecoder(r.Body).Decode(&reg)
	err := s.Store.SendDataForRegistration(reg)
	if err != nil {
		w.Write([]byte(err.Error()))
	}else{
		w.Write([]byte("User registration was successful!!"))
	}
}
