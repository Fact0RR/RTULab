package store

type LernStruct struct {
	Id           int
	Transport    string
	CameraId    string
	ViolationId string
	ViolationValue string
	Type string
}

func (s * Store) GetLernFromDB(page int, elem int, skill int, cam_type string, time_start, time_end string) ([]LernStruct,error){
	var ls []LernStruct
	rows, err := s.DB.Query("select id,transport,camera_id, violation_id,violation_value,type  from get_lern_with_interval($1,$2,$3,$4,$5,$6)", page, elem, skill,cam_type,time_start,time_end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := LernStruct{}
		err := rows.Scan(&n.Id, &n.Transport, &n.CameraId, &n.ViolationId,&n.ViolationValue,&n.Type)
		if err != nil {
			return nil, err
		}
		ls = append(ls, n)
	}

	return ls, nil
}