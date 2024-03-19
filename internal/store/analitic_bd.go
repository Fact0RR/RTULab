package store

import "github.com/Fact0RR/RTULab/internal"

func (s *Store) GetAnaliticFromDB(id int, start, end string) ([]internal.AnaliticFromDB, error) {
	arrNonVer := make([]internal.AnaliticFromDB, 0)

	rows, err := s.DB.Query("select id,employee_id, is_correct,date from get_employee_analitic_with_interval($1,$2,$3)", id, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := internal.AnaliticFromDB{}
		err := rows.Scan(&n.Id, &n.EmployeeId, &n.IsCorrect, &n.Date)
		if err != nil {
			return nil, err
		}
		arrNonVer = append(arrNonVer, n)
	}

	return arrNonVer, nil
}

func (s *Store) GetCountUnkownFromDB(id int) (int, error) {
	var count int
	err := s.DB.QueryRow("select count(*) from excesses_employees_pool where isviolation is null and employee_id = $1", id).Scan(&count)
	if err!=nil{
		return 0,err
	}
	return count, nil
}
