package excel

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ViolationsFine struct {
	Violation_id string
	Fine         int
}

func GetViolationsFromExcel(path string) []ViolationsFine{
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

	arr_vf := make([]ViolationsFine, 0,len(rows))

	fmt.Println(len(arr_vf))
	for _, row := range rows {

		for i, colCell := range row {
			if i%2 == 0 {
				arr_vf = append(arr_vf, ViolationsFine{
					Violation_id: colCell,
				})
			} else {
				r, err := strconv.Atoi(colCell)
				if err != nil {
					panic(err)
				}
				arr_vf[len(arr_vf)-1].Fine = r
			}
		}
	}
	return arr_vf
	
}

func GetMapViolationsFineFromExcel(path string) map[string]string{
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rows, err := f.GetRows("Лист1")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	map_vf := make(map[string]string)

	
	for _, row := range rows {
		map_vf[row[0]] = row[1]
	}

	return map_vf
}