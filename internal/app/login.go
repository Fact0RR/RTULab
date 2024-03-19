package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Fact0RR/RTULab/internal"
	"github.com/Fact0RR/RTULab/internal/store"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login internal.LoginStruct
	json.NewDecoder(r.Body).Decode(&login)
	check,data, err := s.Store.CheckUser(login)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if check {
		refreshToken := getTokenJWT(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, data)
		cookieRefresh := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		accessToken := getTokenJWT(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, data)
		cookieAccess := http.Cookie{
			Name:     "accessToken",
			Value:    accessToken,
			Path:     "/",
			MaxAge:   s.Conf.SecretAccessKeyLifeInHoures,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookieRefresh)
		http.SetCookie(w, &cookieAccess)

		w.Write([]byte("welcome " + login.Login + "!!"))
	} else {
		w.Write([]byte("Wrong login or password!!"))
	}
}

func getTokenJWTforAdmin(key string, life int, data store.AdminData) string{
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = data.Id
	claims["name"] = data.Name
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(life)).Unix()

	tokeString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	return tokeString
}

func getTokenJWT(key string, life int, data store.UserData) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = data.Id
	claims["skill"] = data.Skill
	claims["verified"] = data.Verified
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(life)).Unix()

	tokeString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	return tokeString
}

func (s *Server) LoginMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		access, refresh, data := WorkWithTokens(w, r, s.Conf.SecretAccessKey, s.Conf.SecretRefreshKey)
		if access {
			next(w, r)
		} else if refresh {
			
			refreshToken := getTokenJWT(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, data)
			cookieRefresh := http.Cookie{
				Name:     "refreshToken",
				Value:    refreshToken,
				Path:     "/",
				MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			accessToken := getTokenJWT(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, data)
			cookieAccess := http.Cookie{
				Name:     "accessToken",
				Value:    accessToken,
				Path:     "/",
				MaxAge:   s.Conf.SecretAccessKeyLifeInHoures,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookieRefresh)
			http.SetCookie(w, &cookieAccess)
			
			next(w,r)
		} else {
			http.Error(w, "Not Authorized", http.StatusBadRequest)
		}

	}
}




func (s *Server) VerifiedMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessCookie, err := r.Cookie("accessToken")
		if err != nil{
			refreshCookie, err := r.Cookie("refreshToken")
			if err != nil {
				log.Println("No refresh token")
				return 
			}

			token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(s.Conf.SecretRefreshKey), nil
			})
			if err != nil {
				log.Println("Bad refresh and acces token")
				return 
			}
			claims, _  := token.Claims.(jwt.MapClaims)
			userData := store.UserData{
				Id: int(claims["id"].((float64))),
				Skill: int(claims["skill"].(float64)),
				Verified: claims["verified"].(bool),
			}
			
			answ,err := s.Store.IsUserVerified(userData.Id)
			if err!=nil{
				w.Write([]byte(err.Error()))
			}
			if answ{
				next(w,r)
			}else{
				w.Write([]byte("User not verified, contact the administration"))
			}
			return
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
		claims, _  := token.Claims.(jwt.MapClaims)
		

		userData := store.UserData{
			Id: int(claims["id"].((float64))),
			Skill: int(claims["skill"].(float64)),
			Verified: claims["verified"].(bool),
		}
		
		answ,err := s.Store.IsUserVerified(userData.Id)
		if err!=nil{
			w.Write([]byte(err.Error()))
		}
		if answ{
			next(w,r)
		}else{
			w.Write([]byte("User not verified, contact the administration"))
		}
	}
}
