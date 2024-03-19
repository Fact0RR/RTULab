package internal

type DirtyAnalitic struct {
	Analitics []AnaliticFromDB
}

type AnaliticFromDB struct {
	Id         int
	EmployeeId int
	IsCorrect  bool
	Date       string
}

type CleanAnalitic struct {
	CountAnswers  int `json:"count_answer"`
	CountCorrect  int `json:"count_correct"`
	CountUnknown  int `json:"count_unknown"`
	CountMaxScore int `json:"count_max_score"`
}
