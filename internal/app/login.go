package app

import (
	"encoding/json"
	"errors"
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
		cookie := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookie)
		s.Store.TokenRefreshMap[refreshToken] = login.Login
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

		cookie, err := r.Cookie("refreshToken")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				
				http.Error(w, "Token not found, please login...", http.StatusBadRequest)
			default:
				log.Println(err)
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}

		

		if len(s.Store.TokenRefreshMap[cookie.Value])>0 {
			next(w, r)
		}else{
			http.Error(w, "Wrong token, please re-login...", http.StatusBadRequest)
			delete(s.Store.TokenRefreshMap, cookie.Value);
		}
		
	}
}
