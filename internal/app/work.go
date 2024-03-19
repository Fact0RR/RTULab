package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Fact0RR/RTULab/internal/model/fine"
	"github.com/Fact0RR/RTULab/internal/store"
	"github.com/golang-jwt/jwt"
)

type HandleExcess struct {
	IdExcess    int  `json:"id_excess"`
	IsViolation bool `json:"is_violation"`
}

func (s *Server) SetAnswerOfExcess(w http.ResponseWriter, r *http.Request) {
	ud := s.getClaims(r)
	var he HandleExcess
	json.NewDecoder(r.Body).Decode(&he)
	err := s.Store.SendAnswerOnExcess(ud.Id, he.IdExcess, he.IsViolation)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	check,err := s.Store.CheckSolve(he.IdExcess)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if check {
		fineData,err := s.Store.GetExcessesForCitizen(he.IdExcess)
		if err != nil{
			log.Fatal(err)
		}
		err = fine.SendFine(fineData,s.Conf,s.Store.ViolationsFine,s.Store.CitizenConnectionData)
		if err != nil{
			log.Fatal(err)
		}
	}

	w.Write([]byte("Answer send on server"))
}

func (s *Server) GetExcessesFromPool(w http.ResponseWriter, r *http.Request) {
	ud := s.getClaims(r)
	workArr, err := s.Store.GetExcessesPoolForEmployee(ud.Id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(workArr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func (s *Server) getClaims(r *http.Request) *store.UserData {
	accessCookie, err := r.Cookie("accessToken")
	if err != nil {
		refreshCookie, err := r.Cookie("refreshToken")
		if err != nil {
			log.Println("No refresh token")
			return nil
		}

		token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(s.Conf.SecretRefreshKey), nil
		})
		if err != nil {
			log.Println("Bad refresh and acces token")
			return nil
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		userData := store.UserData{
			Id:       int(claims["id"].((float64))),
			Skill:    int(claims["skill"].(float64)),
			Verified: claims["verified"].(bool),
		}

		return &userData
	}
	token, err := jwt.Parse(accessCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(s.Conf.SecretAccessKey), nil
	})
	if err != nil {
		log.Println("Bad access token")
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	userData := store.UserData{
		Id:       int(claims["id"].((float64))),
		Skill:    int(claims["skill"].(float64)),
		Verified: claims["verified"].(bool),
	}

	return &userData
}
