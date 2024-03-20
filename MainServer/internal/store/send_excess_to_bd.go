package store

import "github.com/Fact0RR/RTULab/internal"

func (s *Store) SendExcessToDB(uc *internal.UnionCamera) error {
	_, err := s.DB.Exec("call createExcessesPool($1,$2,$3,$4,$5,$6,$7)", uc.Transport, uc.CameraID, uc.ViolationID,uc.ViolationValue,uc.Skill,uc.DateTime,uc.Photo)
	if err != nil {
		return err
	}
	return nil
}