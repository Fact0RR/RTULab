package store

type CurrentStatistic struct {
	Excess_id int `json:"excess_id"`
	Login     string `json:"login"`
	Answer    bool `json:"answer"`
	Skill 	  int `json:"skill"`
}

func (s *Store) GetCurrentStatisticfromBD() ([]CurrentStatistic, error) {
	arrNonVer := make([]CurrentStatistic, 0)

	rows, err := s.DB.Query("select ex.id,em.login,eep.isViolation,em.skill from excesses_employees_pool eep JOIN excesses ex ON ex.id = eep.excess_id JOIN employees em ON em.id = eep.employee_id WHERE eep.isviolation IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := CurrentStatistic{}
		err := rows.Scan(&n.Excess_id, &n.Login, &n.Answer,&n.Skill)
		if err != nil {
			return nil, err
		}
		arrNonVer = append(arrNonVer, n)
	}

	return arrNonVer, nil
}
