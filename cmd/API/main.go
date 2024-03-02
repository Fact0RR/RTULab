package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

import "github.com/Fact0RR/RTULab/internal/app"

func main() {

	server := app.New()
	server.StartApp()	

	// mySigningKey := "sss"

	// tokenString := getTokenJWT(mySigningKey,5,"lll")
	// for{
	// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("There was an error")
	// 		}
	// 		return []byte(mySigningKey), nil
	// 	})
	// 	if err != nil {
	// 		log.Println("12312")
	// 		log.Println(err.Error())
	// 	}else{
	// 		if token.Valid{
	// 			log.Println("tokenValid")
	// 		}else{
	// 			log.Println("tokenInValid")
	// 		}
	// 	}
	// 	time.Sleep(time.Microsecond)
	//}
}


// func getTokenJWT(key string, life int, login string) string {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["user"] = login
// 	claims["exp"] = time.Now().Add(time.Second * time.Duration(life)).Unix()

// 	tokeString, err := token.SignedString([]byte(key))
// 	if err != nil {
// 		log.Fatal("sdasdadsa", err)
// 	}
// 	return tokeString
// }