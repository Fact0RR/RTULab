package analitic

import "github.com/Fact0RR/RTULab/internal"

func GetCleanAnalitic(dirty []internal.AnaliticFromDB, unknown int) internal.CleanAnalitic{
	var cA internal.CleanAnalitic

	cA.CountAnswers = len(dirty)
	cA.CountUnknown = unknown

	maxScore:=0
	score:=0
	countCorrect:=0

	for _,an :=range dirty{
		if an.IsCorrect{
			countCorrect++
			score++
			if maxScore<score{
				maxScore = score
			}
		}else{
			score = 0
		}

	}
	cA.CountMaxScore = maxScore
	cA.CountCorrect = countCorrect
	return cA
}