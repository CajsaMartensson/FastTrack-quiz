package main

import (
	"encoding/json"
	"fasttrack-quiz/server"
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("--Välkommen till quizet--")

	//Här registrerar endpoints
	http.HandleFunc("/questions", getQuestions)
	http.HandleFunc("/submitAnswers", submitAnswers)

	var serverError = http.ListenAndServe(":8080", nil)

	if serverError != nil{
		fmt.Println("kunde inte starta serven")
	}
}

func getQuestions (responseWriter http.ResponseWriter, httpRequest *http.Request){
	json.NewEncoder(responseWriter).Encode(server.QuestionList)
}

func submitAnswers(reponseWriter http.ResponseWriter, httpRequest *http.Request){
	if httpRequest.Method != http.MethodPost{
		http.Error(reponseWriter, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var answers []int
	json.NewDecoder(httpRequest.Body).Decode(&answers)

	var correctCount = 0
	for i, selectedAnswer := range answers{
		if i < len(server.QuestionList){
			chosenNumber := server.QuestionList[i].Options[selectedAnswer]
			if chosenNumber == server.QuestionList[i].CorrectAnswer{
				correctCount++
			}
		}
	}

	var betterThan = 0
	for _, score := range server.AllScores{
		if correctCount >= score{
			betterThan++
		}
	}

	var percent = 0.0
	if len(server.AllScores) > 0 {
		percent = (float64(betterThan) / float64(len(server.AllScores))) * 100
	} else {
		percent = 100.0
	}
	server.AllScores = append(server.AllScores, correctCount)

	var response = server.QuizResponse{
		Score: correctCount,
		Percent: percent,
	}

	json.NewEncoder(reponseWriter).Encode(response);
}