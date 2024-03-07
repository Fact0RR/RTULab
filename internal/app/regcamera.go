package app

import (
	"encoding/json"
	"net/http"

	"github.com/Fact0RR/RTULab/internal"
)

func (s *Server) RegCamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cam internal.Camera
	json.NewDecoder(r.Body).Decode(&cam)
	id,err := s.Store.SendDataForRegistrationCamera(cam)
	if err != nil{
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(id))
}