package store

import "github.com/Fact0RR/RTULab/internal"

type UserData struct {
	Id       int
	Skill    int
	Verified bool
}

type AdminData struct{
	Id int
	Name string
}

func (s *Store) CheckUser(login internal.LoginStruct) (bool, UserData, error) {
	res := s.DB.QueryRow("SELECT id, skill, verified from employees where password = crypt($1, password) and login = $2", login.Password, login.Login)
	data := UserData{}
	err := res.Scan(&data.Id,&data.Skill,&data.Verified)
	if err != nil {
		return false, UserData{}, err
	}
	
	return true, data ,nil
}

func (s *Store) CheckAdmin(login internal.LoginStruct) (bool, *AdminData, error) {	
res := s.DB.QueryRow("SELECT id, name from masters where password = crypt($1, password) and login = $2", login.Password, login.Login)
data := AdminData{}
err := res.Scan(&data.Id,&data.Name)
if err != nil {
	return false, &AdminData{}, err
}
return true, &data ,nil
}