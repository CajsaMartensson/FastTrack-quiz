package server

type Question struct {
	ID            int		`json:"id"`
	Text          string	`json:"text"`
	Options       []int		`json:"options"`
	CorrectAnswer int
}

var QuestionList = []Question{
	{ID: 1, Text: "Vad är 1 + 1?", Options: []int{1, 2, 3}, CorrectAnswer: 2},
    {ID: 2, Text: "Vad är 10 - 5?", Options: []int{5, 15, 20}, CorrectAnswer: 5},
    {ID: 3, Text: "Vad är 6 * 7?", Options: []int{36, 42, 48, 54}, CorrectAnswer: 42},
    {ID: 4, Text: "Vad är roten ur 81?", Options: []int{7, 8, 9, 10}, CorrectAnswer: 9},
    {ID: 5, Text: "Om x + 5 = 12, vad är x?", Options: []int{5, 6, 7, 8}, CorrectAnswer: 7},
}

type QuizResponse struct {
	Score 		int			`json:"score"`
	Percent 	float64		`json:"percent"`
}

var AllScores = []int{}