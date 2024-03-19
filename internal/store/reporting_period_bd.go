package store

func (s *Store) CallStopReportingPeriod()error{
	_,err:= s.DB.Exec("call endOfReportingPeriod();")
	if err!=nil{
		return err
	}
	return nil
}