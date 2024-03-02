package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Fact0RR/RTULab/internal"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login internal.LoginStruct
	json.NewDecoder(r.Body).Decode(&login)
	check, err := s.Store.CheckUser(login)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if check {
		refreshToken := getTokenJWT(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, login.Login)
		cookieRefresh := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		accessToken := getTokenJWT(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, login.Login)
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

func getTokenJWT(key string, life int, login string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = login
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(life)).Unix()

	tokeString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal("sdasdadsa", err)
	}
	return tokeString
}

func (s *Server) LoginMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		access, refresh := WorkWithTokens(w, r, s.Conf.SecretAccessKey, s.Conf.SecretRefreshKey)
		if access {
			next(w, r)
		} else if refresh {
			
			refreshToken := getTokenJWT(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, "refresh")
			cookieRefresh := http.Cookie{
				Name:     "refreshToken",
				Value:    refreshToken,
				Path:     "/",
				MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			accessToken := getTokenJWT(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, "access")
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
