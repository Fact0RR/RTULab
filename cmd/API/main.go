package main

import (
	"net/http"

	"github.com/Fact0RR/RTULab/internal/app"
)

func main() {
	//конструктор сервиса
	server := app.New()
	//проверка работы вспомогательного сервера (он отправляет и записывает фотографии)
	r, err := http.Get(server.Conf.Photo_Server_URL+"123")
	if err != nil {
		panic(err)
	}
	r.Body.Close()
	//запуск сервиса
	server.StartApp()	
	
}