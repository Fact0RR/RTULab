package store

type Work struct{
	Id int `json:"id"`
	Transport string `json:"transport"`
	Camera_id string `json:"camera_id"`
	Violation_id string `json:"violation_id"`
	Violation_value string `json:"violation_value"`
	Skill int `json:"skill"`
	Time string `json:"time"`
	Photo string `json:"photo"`
}
type CitizenFine struct{
	Id int `json:"id"`
	Transport string `json:"transport"`
	Camera_id string `json:"camera_id"`
	Violation_id string `json:"violation_id"`
	Violation_value string `json:"violation_value"`
	Skill int `json:"skill"`
	Time string `json:"time"`
	Photo string `json:"photo"`
	CoordinateX float64 `json:"coordinateX"`
	CoordinateY float64 `json:"coordinateY"`
}

func (s *Store) GetExcessesPoolForEmployee(id int) ([]Work,error){
	arr_work := make([]Work,0)

	rows, err := s.DB.Query("select e.id,e.transport, e.camera_id, e.violation_id,e.violation_value,e.skill,e.datetime, e.photo from excesses_employees_pool eep join excesses e on e.id = eep.excess_id where eep.employee_id = $1 and eep.isviolation is null", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := Work{}
		err := rows.Scan(&n.Id, &n.Transport, &n.Camera_id,&n.Violation_id,&n.Violation_value,&n.Skill,&n.Time,&n.Photo)
		if err != nil {
			return nil, err
		}
		arr_work = append(arr_work, n)
	}

	return arr_work, nil
}

func (s *Store) GetExcessesForCitizen(id int) ([]CitizenFine,error){
	arr_work := make([]CitizenFine,0)

	rows, err := s.DB.Query("select e.id,e.transport, e.camera_id, e.violation_id,e.violation_value,e.skill,e.datetime, e.photo, c.coordinatex, c.coordinatey from excesses e join cameras c on c.id= e.camera_id  where e.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := CitizenFine{}
		err := rows.Scan(&n.Id, &n.Transport, &n.Camera_id,&n.Violation_id,&n.Violation_value,&n.Skill,&n.Time,&n.Photo,&n.CoordinateX,&n.CoordinateY)
		if err != nil {
			return nil, err
		}
		arr_work = append(arr_work, n)
	}

	return arr_work, nil
}

func (s *Store) SendAnswerOnExcess(emp_id int, exc_id int, answ bool) error{
	
	_,err:= s.DB.Exec("update excesses_employees_pool set isviolation = $1 where excess_id = $2 and employee_id = $3",answ,exc_id,emp_id)
	return err
}

func (s*Store) CheckSolve(id_exc int) (bool,error){
	var res bool
	err:= s.DB.QueryRow("select check_excess_on_readiness($1)",id_exc).Scan(&res)
	return res,err
}