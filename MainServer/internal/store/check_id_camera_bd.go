package store

func (s *Store) CheckCameraID(id string) bool {
	var count int
	s.DB.QueryRow("select count(*) from cameras where id = $1",id).Scan(&count)
	return count == 1
}
