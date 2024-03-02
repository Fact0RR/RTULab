package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func WorkWithTokens(w http.ResponseWriter, r *http.Request, secretAccess, secretRefresh string) (bool, bool) {
	accessCookie, err := r.Cookie("accessToken")
	if err != nil {
		log.Println("No access token")
	}else{
		token, err := jwt.Parse(accessCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secretAccess), nil
		})
		if err != nil {
			log.Println("Bad access token")
		}
		if token.Valid {
			return true, false
		}
	}

	refreshCookie, err := r.Cookie("refreshToken")
	if err != nil {
		log.Println("No refresh token")
		return false, false
	}

	token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secretRefresh), nil
	})
	if err != nil {
		log.Println("Bad refresh and acces token")
		return false, false
	}
	if token.Valid {
		log.Println("refresh token valid")
		return false,true
	}
	log.Println("all invalid")
	return false,false
}
