package main

type Questions struct {
	ID            int
	Text          string
	Options       []int
	CorrectAnswer int
}

func main() {

	var questionsList = []Questions{
		{ID: 1, Text: "What is 1 + 1", Options: []int{1, 2, 3}, CorrectAnswer: 2},
		{ID: 2, Text: "What is 10 - 5", Options: []int{5, 15, 20}, CorrectAnswer: 5},
	}
}