package app

import (
	"encoding/json"
	"net/http"

	"github.com/Fact0RR/RTULab/internal"
)

func (s *Server) LoginAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login internal.LoginStruct
	json.NewDecoder(r.Body).Decode(&login)
	check,data, err := s.Store.CheckAdmin(login)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if check {
		refreshToken := getTokenJWTforAdmin(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, *data)
		cookieRefresh := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		accessToken := getTokenJWTforAdmin(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, *data)
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

func (s *Server) LoginAdminMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		access, refresh, data := WorkWithAdminTokens(w, r, s.Conf.SecretAccessKey, s.Conf.SecretRefreshKey)
		if access {
			next(w, r)
		} else if refresh {
			
			refreshToken := getTokenJWTforAdmin(s.Conf.SecretRefreshKey, s.Conf.SecretRefresKeyLifeInHoures, data)
			cookieRefresh := http.Cookie{
				Name:     "refreshToken",
				Value:    refreshToken,
				Path:     "/",
				MaxAge:   s.Conf.SecretRefresKeyLifeInHoures,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			accessToken := getTokenJWTforAdmin(s.Conf.SecretAccessKey, s.Conf.SecretAccessKeyLifeInHoures, data)
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