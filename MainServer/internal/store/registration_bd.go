package store

import (
	"github.com/Fact0RR/RTULab/internal"
)

func (s *Store) SendDataForRegistration(reg internal.RegStruct) error {

	_, err := s.DB.Exec("call createemployee($1,$2,$3,$4,$5,$6)", reg.Name, reg.Email, reg.Skill,reg.Login,reg.Password,reg.Photob64)
	if err != nil {
		return err
	}
	return nil
}
