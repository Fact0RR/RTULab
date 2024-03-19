package store

import "github.com/Fact0RR/RTULab/internal"

func (s *Store) IsUserVerified(id int) (bool, error) {
	var verified bool
	err := s.DB.QueryRow("select verified from employees where id = $1", id).Scan(&verified)
	if err != nil {
		return false, err
	}
	return verified, nil
}

func (s *Store) GetNonVerified() ([]internal.NonVerify, error) {
	arrNonVer := make([]internal.NonVerify, 0)

	rows, err := s.DB.Query("select id,name,login,skill from employees where verified = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := internal.NonVerify{}
		err := rows.Scan(&n.Id, &n.Name, &n.Login, &n.Skill)
		if err != nil {
			return nil, err
		}
		arrNonVer = append(arrNonVer, n)
	}

	return arrNonVer, nil
}

func (s *Store) SetVerify(login string, verify bool) (int, error) {
	var id int
	err:= s.DB.QueryRow("update employees set verified = $1 where verified = $2 returning id",verify,
		login).Scan(&id)
	if err!=nil{
		return 0,err
	}
	return id, nil
}
