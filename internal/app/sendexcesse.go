package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Fact0RR/RTULab/internal"
	"github.com/Fact0RR/RTULab/internal/model/validation"
)

type cameraType struct {
	CameraType string `json:"camera_type"`
}

func (s *Server) SendExcess(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	data, err := deserData(&body)
	if err != nil {
		log.Fatal(err)
	}
	var ct cameraType
	json.Unmarshal(data, &ct)
	var uc *internal.UnionCamera
	switch ct.CameraType {
	case "camerus1":
		var camera internal.Camera1
		json.Unmarshal(data, &camera)
		uc = camera.ToUnionStruct()
	case "camerus2":
		var camera internal.Camera2
		json.Unmarshal(data, &camera)
		uc = camera.ToUnionStruct()
	case "camerus3":
		var camera internal.Camera3
		json.Unmarshal(data, &camera)
		uc = camera.ToUnionStruct()
	default:
		log.Println("Неизвестный тип камеры")
		log.Fatal("Неизвестный тип камеры")
	}

	if !s.Store.CheckCameraID(uc.CameraID) {
		log.Fatal("Id камеры не найден")
	}

	log.Println(s.Store.ViolationsFine[uc.ViolationID])

	if len(s.Store.ViolationsFine[uc.ViolationID])==0{
		log.Fatal("Id нарушения не найден")
	}

	err = validation.CheckCameraData(uc)
	if err != nil{
		log.Fatal(err)
	}


	//err = s.Store.SendExcessToDB(uc)
	//if err != nil{
	//	log.Fatal(err)
	//}
	w.Write([]byte("Excess send"))

}

func deserData(by *[]byte) ([]byte, error) {

	r := bytes.NewReader(*by)
	resp, err := http.Post("https://recruit.rtuitlab.dev/deserialize", "application/json", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
