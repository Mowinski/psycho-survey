package herzberg



type SurveyTestItem struct {
	SelectedFirst bool
	
	firstLetter rune
	secondLetter rune
}

type MotivationHerzbergTestAnswers struct {
	Answers []SurveyTestItem `json:"answers"`
}


type SummaryLatters struct {
	Aanswer int8
	Banswer int8
	Canswer int8
	Danswer int8
	Eanswer int8
	Fanswer int8
	Ganswer int8
	Hanswer int8
}
