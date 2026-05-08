package cmd

import (
	"bytes"
	"encoding/json"
	"fasttrack-quiz/server"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quiz called")

		resp, err:= http.Get("http://localhost:8080/questions")
		if err != nil{
			fmt.Println("Kunde inte nå API:et")
			return
		}
		defer resp.Body.Close()

		
		var questions []server.Question

		err = json.NewDecoder(resp.Body).Decode(&questions)

		if err != nil{
			fmt.Println("Kunde inte läsa frågorna")
			return
		}

		runQuiz(questions)
	},
}

func runQuiz(questions []server.Question){
	var userAnswers []int

	for i, q := range questions{
		fmt.Printf("\nFråga %d: %s\n", i+1, q.Text)
	
		for j, option := range q.Options{
			fmt.Printf("%d) %d\n", j+1, option)
		}

		fmt.Print("Ditt svar (nummer): ")
		var userAnswer int
	
		fmt.Scanln(&userAnswer)

		userAnswer = userAnswer - 1

		userAnswers = append(userAnswers, userAnswer)
	}
	postAnswers(userAnswers)
}

func postAnswers(answers[] int){
	//gör om listan till json-format
	jsonData, err := json.Marshal(answers)
	
	if err != nil {
    fmt.Println("Fel vid skapande av JSON")
    return
}

	resp, err := http.Post(
		"http://localhost:8080/submitAnswers",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	defer resp.Body.Close()

	var result server.QuizResponse

	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("\n--- QUIZ KLART ---\n")
	fmt.Printf("Du fick %d rätt!\n", result.Score)
	fmt.Printf("Du var bättre än %.1f%% av alla andra.\n", result.Percent)
}

func init() {
	rootCmd.AddCommand(quizCmd)
}