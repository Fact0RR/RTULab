package store

import "github.com/Fact0RR/RTULab/internal"

func (s *Store) CheckUser(login internal.LoginStruct) (bool, error) {
	res, err := s.DB.Query("SELECT id from employees where password = crypt($1, password) and login = $2", login.Password, login.Login)
	if err != nil {
		res.Close()
		return false, err
	}
	defer res.Close()
	if res.Next() {
		return true, err
	}
	return false, nil
}
