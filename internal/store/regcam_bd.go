package store

import "github.com/Fact0RR/RTULab/internal"

func (s *Store) SendDataForRegistrationCamera(reg internal.Camera) (string, error) {
	var id string
	err := s.DB.QueryRow("insert into cameras (id, type, coordinateX, coordinateY, description) values ($1,$2,$3,$4,$5) returning id",
		reg.Id, reg.Type, reg.Xcoordinate, reg.Ycoordinate, reg.Description).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
