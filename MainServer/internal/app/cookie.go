package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fact0RR/RTULab/internal/store"
	"github.com/golang-jwt/jwt"
)

func WorkWithTokens(w http.ResponseWriter, r *http.Request, secretAccess, secretRefresh string) (bool, bool, store.UserData) {
	accessCookie, err := r.Cookie("accessToken")
	if err != nil {
		log.Println("No access token")
	} else {
		token, err := jwt.Parse(accessCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secretAccess), nil
		})
		if err != nil {
			log.Println("Bad access token")
		}
		claims, ok  := token.Claims.(jwt.MapClaims)
		if token.Valid && ok{

			return true, false, store.UserData{
				Id: int(claims["id"].((float64))),
				Skill: int(claims["skill"].(float64)),
				Verified: claims["verified"].(bool),
			}
		}
	}

	refreshCookie, err := r.Cookie("refreshToken")
	if err != nil {
		log.Println("No refresh token")
		return false, false, store.UserData{}
	}

	token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secretRefresh), nil
	})
	if err != nil {
		log.Println("Bad refresh and acces token")
		return false, false,store.UserData{}
	}
	claims, ok  := token.Claims.(jwt.MapClaims)
	if token.Valid && ok{
		log.Println("refresh token valid")
		return false, true,store.UserData{
			Id: int(claims["id"].(float64)),
			Skill: int(claims["skill"].(float64)),
			Verified: claims["verified"].(bool),
		}
	}
	log.Println("all invalid")
	return false, false, store.UserData{}
}



func WorkWithAdminTokens(w http.ResponseWriter, r *http.Request, secretAccess, secretRefresh string) (bool, bool, store.AdminData) {
	accessCookie, err := r.Cookie("accessToken")
	if err != nil {
		log.Println("No access token")
	} else {
		token, err := jwt.Parse(accessCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secretAccess), nil
		})
		if err != nil {
			log.Println("Bad access token")
		}
		claims, ok  := token.Claims.(jwt.MapClaims)
		if token.Valid && ok{

			return true, false, store.AdminData{
				Id: int(claims["id"].((float64))),
				Name: claims["name"].(string),
			}
		}
	}

	refreshCookie, err := r.Cookie("refreshToken")
	if err != nil {
		log.Println("No refresh token")
		return false, false, store.AdminData{}
	}

	token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(secretRefresh), nil
	})
	if err != nil {
		log.Println("Bad refresh and acces token")
		return false, false,store.AdminData{}
	}
	claims, ok  := token.Claims.(jwt.MapClaims)
	if token.Valid && ok{
		log.Println("refresh token valid")
		return false, true,store.AdminData{
			Id: int(claims["id"].((float64))),
			Name: claims["name"].(string),
		}
	}
	log.Println("all invalid")
	return false, false, store.AdminData{}
}
