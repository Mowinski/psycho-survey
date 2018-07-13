package herzberg_test

import (
	"testing"
	"bytes"
	"fmt"
	
	"github.com/Mowinski/psycho-survey/herzberg"
)


func assertAnswer(t *testing.T, number int8, answers herzberg.MotivationHerzbergTestAnswers, expected bool) {
	if answers.Answers[number].SelectedFirst != expected {
		t.Errorf(
			"Answer number %d is incorect, expected %v got %v",
			number,
			expected,
			answers.Answers[number].SelectedFirst,
		)
	}
}

func assertSummary(t *testing.T, summary herzberg.SummaryLatters, answers [8]int8) {
	if summary.Aanswer != answers[0] {
		t.Errorf("Wrong 'A' answer got %d, expected %d", summary.Aanswer, answers[0])
	}

	if summary.Banswer != answers[1] {
		t.Errorf("Wrong 'B' answer got %d, expected %d", summary.Banswer, answers[1])
	}

	if summary.Canswer != answers[2] {
		t.Errorf("Wrong 'C' answer got %d, expected %d", summary.Canswer, answers[2])
	}

	if summary.Danswer != answers[3] {
		t.Errorf("Wrong 'D' answer got %d, expected %d", summary.Danswer, answers[3])
	}

	if summary.Eanswer != answers[4] {
		t.Errorf("Wrong 'E' answer got %d, expected %d", summary.Eanswer, answers[4])
	}

	if summary.Fanswer != answers[5] {
		t.Errorf("Wrong 'F' answer got %d, expected %d", summary.Fanswer, answers[5])
	}

	if summary.Ganswer != answers[6] {
		t.Errorf("Wrong 'G' answer got %d, expected %d", summary.Ganswer, answers[6])
	}

	if summary.Hanswer != answers[7] {
		t.Errorf("Wrong 'H' answer got %d, expected %d", summary.Hanswer, answers[7])
	}
}
func TestGetLetterWhenSelectedFirstAnswer(t *testing.T) {
	testItem := herzberg.NewTestItem(true, 'A', 'B')
	
	result := testItem.GetLetter()

	if result != 'A' {
		t.Errorf("Wrong returned letter, expected 'A' got '%c'", result)
	}
}

func TestGetLetterWhenSelectedSecondAnswer(t *testing.T) {
	testItem := herzberg.NewTestItem(false, 'A', 'B')
	
	result := testItem.GetLetter()

	if result != 'B' {
		t.Errorf("Wrong returned letter, expected 'B' got '%c'", result)
	}
}

func TestReadFromJSON(t *testing.T) {
	valid_answers := `
	{
		"answers": [
			true, false, true, true, false, true, false, true, false,
			true, false, true, false, true, true, false, true, true,
			false, true, true, false, true, false, true, true, false,
			true
		]
	}
	`
	r := bytes.NewBufferString(valid_answers)
	answers, err := herzberg.Read(r)

	if err != nil {
		t.Fatalf("Error occurs %s", err)
	}
	fmt.Println(answers.Answers[0])
	assertAnswer(t, 0, answers, true)
	assertAnswer(t, 1, answers, false)
	assertAnswer(t, 2, answers, true)
	assertAnswer(t, 3, answers, true)
	assertAnswer(t, 4, answers, false)
	assertAnswer(t, 5, answers, true)
	assertAnswer(t, 6, answers, false)
	assertAnswer(t, 7, answers, true)
	assertAnswer(t, 8, answers, false)
	assertAnswer(t, 9, answers, true)
	assertAnswer(t, 10, answers, false)
	assertAnswer(t, 11, answers, true)
	assertAnswer(t, 12, answers, false)
	assertAnswer(t, 13, answers, true)
	assertAnswer(t, 14, answers, true)
	assertAnswer(t, 15, answers, false)
	assertAnswer(t, 16, answers, true)
	assertAnswer(t, 17, answers, true)
	assertAnswer(t, 18, answers, false)
	assertAnswer(t, 19, answers, true)
	assertAnswer(t, 20, answers, true)
	assertAnswer(t, 21, answers, false)
	assertAnswer(t, 22, answers, true)
	assertAnswer(t, 23, answers, false)
	assertAnswer(t, 24, answers, true)
	assertAnswer(t, 25, answers, true)
	assertAnswer(t, 26, answers, false)
	assertAnswer(t, 27, answers, true)
}


func TestReadFromJSONWhenToSmallAnswers(t *testing.T) {
	valid_answers := `
	{
		"answers": [
			true, false, true, true, false, true, false, true, false,
			true, false, true, false, true, true, false, true, true,
			false, true, true, false, true, false, true, true, false
		]
	}
	`
	r := bytes.NewBufferString(valid_answers)
	_, err := herzberg.Read(r)

	if err == nil || err.Error() != "Wrong number of answer. Expected 28" {
		t.Fatalf("Error do not occure, %v", err)
	}
}


func TestReadFromJSONWhenToMuchAnswers(t *testing.T) {
	valid_answers := `
	{
		"answers": [
			true, false, true, true, false, true, false, true, false,
			true, false, true, false, true, true, false, true, true,
			false, true, true, false, true, false, true, true, false,
			true, true
		]
	}
	`
	r := bytes.NewBufferString(valid_answers)
	_, err := herzberg.Read(r)

	if err == nil || err.Error() != "Wrong number of answer. Expected 28" {
		t.Fatalf("Error do not occure, %v", err)
	}
}


func TestCollectSummaryWhenAlwaysSecondOptionIsSelected(t *testing.T) {
	answersItems := []herzberg.SurveyTestItem{
		herzberg.NewTestItem(false, 'A', 'B'),
		herzberg.NewTestItem(false, 'C', 'D'),
		herzberg.NewTestItem(false, 'D', 'B'),
		herzberg.NewTestItem(false, 'A', 'A'),
	}
	expectedAnswers := [8]int8{1, 2, 0, 1, 0, 0, 0, 0}
	answers := herzberg.MotivationHerzbergTestAnswers{
		Answers: answersItems,
	}

	summary, err := answers.GetSummary()
	if err != nil {
		t.Fatal(err)
	}
		
	assertSummary(t, summary, expectedAnswers)
}
