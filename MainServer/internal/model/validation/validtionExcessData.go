package validation

// type UnionCamera struct {
// 	Transport      string
// 	CameraID       string
// 	ViolationID    string
// 	ViolationValue string
// 	Skill          int
// 	DateTime       string
// 	Photo          string
// }

import (
	"errors"
	"regexp"
	"time"

	"github.com/Fact0RR/RTULab/internal"
)

func CheckCameraData(uc *internal.UnionCamera) error{

	
	matched, _ := regexp.MatchString(`^[A-Z]{3}[0-9]{3}([0-9]{2}|[0-9]{3})$`, uc.Transport)
	if !matched {
		return errors.New("номер транспорта не прошел валидацию : " + uc.Transport)
	}
	matched, _ = regexp.MatchString(`^[a-z0-9]{8}\-[a-z0-9]{4}\-[a-z0-9]{4}\-[a-z0-9]{17}$`,uc.CameraID)
	if !matched {
		return errors.New("id камеры не прошел валидацию : "+ uc.CameraID)
	}
	matched, _ = regexp.MatchString(`^[a-z0-9]{8}\-[a-z0-9]{4}\-[a-z0-9]{4}\-[a-z0-9]{4}\-[a-z0-9]{12}$`,uc.ViolationID)
	if !matched {
		return errors.New("id нарушения не прошел валидацию : "+ uc.ViolationID)
	}

	if len(uc.ViolationValue)==0{
		return errors.New("violation_value не прошел валидацию : "+ uc.ViolationValue)
	}

	if uc.Skill<1 || uc.Skill>4{
		return errors.New("skill не прошел валидацию : ")
	}

	err := parse(uc.DateTime)
	if err != nil {
		return errors.New("дата не прошла валидацию: "+err.Error())
	}

	if len(uc.Photo)==0{
		return errors.New("фото случая не прошло валидацию :"+uc.Photo)
	}
	
	return nil
}

func parse(s string) error {
	_, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return  err
	}
	return nil
}