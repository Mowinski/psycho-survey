package herzberg

import (
	"io"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

func (s *SurveyTestItem) UnmarshalJSON(b []byte) (err error) {
	s.SelectedFirst, err = strconv.ParseBool(string(b))	
	return err
}

func (s *SurveyTestItem) GetLetter() (rune) {
	if s.SelectedFirst {
		return s.firstLetter
	}
	return s.secondLetter 
}

func NewTestItem(selectedFirst bool, firstLetter, secondLetter rune) SurveyTestItem {
	item := SurveyTestItem{
		SelectedFirst: selectedFirst,
		firstLetter: firstLetter,
		secondLetter: secondLetter,
	}

	return item
}

func Read(r io.Reader) (MotivationHerzbergTestAnswers, error) {
	var testAnswers MotivationHerzbergTestAnswers
	content, err := ioutil.ReadAll(r)	
	if err != nil {
		return testAnswers, err
	}

	err = json.Unmarshal(content, &testAnswers)	

	if err != nil {
		return testAnswers, err
	}
	
	if len(testAnswers.Answers) != 28 {
		return testAnswers, fmt.Errorf("Wrong number of answer. Expected 28")
	}

	for index, _ := range testAnswers.Answers {
		testAnswers.Answers[index].firstLetter = answersLetters[index][0]
		testAnswers.Answers[index].secondLetter = answersLetters[index][1]
	}

	return testAnswers, nil
}

func (m MotivationHerzbergTestAnswers) GetSummary() (SummaryLatters, error) {
	var summary SummaryLatters

	for _, answer := range m.Answers {
		switch answer.GetLetter() {
		case 'A':
			summary.Aanswer++;
		case 'B':
			summary.Banswer++;
		case 'C':
			summary.Canswer++;
		case 'D':
			summary.Danswer++;
		case 'E':
			summary.Eanswer++;
		case 'F':
			summary.Fanswer++;
		case 'G':
			summary.Ganswer++;
		case 'H':
			summary.Hanswer++;
		}
	}
	return summary, nil
}
