package fine

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Fact0RR/RTULab/config"
	"github.com/Fact0RR/RTULab/internal/model/excel"
	"github.com/Fact0RR/RTULab/internal/store"
)

func SendFine(data []store.CitizenFine, conf *config.Config,finePrices map[string]string, connectionData map[string]excel.Citizen) error{
	finePrice := finePrices[data[0].Violation_id]
	chatTelegram := connectionData[data[0].Transport].Telegram
	log.Println()
	log.Println(data[0].Violation_id)
	d := data[0]
	finalStirng:="Координаты: \n x "+fmt.Sprintf("%.2f", d.CoordinateX)+"\n y "+fmt.Sprintf("%.2f", d.CoordinateY)+"\nТип штрафа и цена: "+finePrice+"\nЗначение: "+d.Violation_value+"\nДата и время: "+d.Time

	

	b,err:=downloadFile(conf.Photo_Server_URL+d.Photo)
	if err != nil {
		return err
	}
	err = SendFineToTelegram(finalStirng, conf.TelegramToken,finePrice,chatTelegram,b)
	if err!=nil{
		return err
	}
	return nil
}

func downloadFile(URL string) ([]byte,error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil,errors.New("Received non 200 response code")
	}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return b,nil
}