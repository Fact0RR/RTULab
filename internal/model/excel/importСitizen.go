package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Citizen struct {
	Transport   string
	Email       string
	Telegram    string
	Vk          string
	PhoneNumber string
}

func GetCitizenFromExcel(path string) []Citizen {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	arr_vf := make([]Citizen, 0, len(rows))

	fmt.Println(len(arr_vf))
	for _, row := range rows {

		for i, colCell := range row {
			if i%5 == 0 {
				arr_vf = append(arr_vf, Citizen{
					Transport: colCell,
				})
			}else if i%5 == 1{
				arr_vf[len(arr_vf)-1].Email = colCell
			} else if i%5 == 2{
				arr_vf[len(arr_vf)-1].Telegram = colCell
			}else if i%5 == 3{
				arr_vf[len(arr_vf)-1].Vk = colCell
			}else if i%5 == 4{
				arr_vf[len(arr_vf)-1].PhoneNumber = colCell
			}
		}
	}
	return arr_vf

}

func GetMapCitizenContactsFromExcel(path string) map[string]Citizen{
	arr := GetCitizenFromExcel(path)
	map_cit := make(map[string]Citizen)
	for _,cit := range arr{
		map_cit[cit.Transport] = cit
	}
	return map_cit
}