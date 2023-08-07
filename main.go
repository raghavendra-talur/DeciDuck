package main

import (
	"fmt"
	"strings"
)

type QuestionFunc func() (string, AnswerMatrix)
type AnswerMatrix map[string]QuestionFunc

var startQuestion, previousQuestion, majorForkQuestion QuestionFunc

func main() {
	startBot()
}

func startBot() {
	startQuestion = Q1
	previousQuestion = nil
	majorForkQuestion = nil

	currentQuestion := startQuestion
	var nextQuestion QuestionFunc

	for {
		questionText, answers := currentQuestion()
		response := askUser(questionText)

		nextQuestion = processResponse(response, answers[response])

		if nextQuestion == nil {
			fmt.Println("Thank you for using the bot. Goodbye!")
			break
		}

		previousQuestion = currentQuestion
		currentQuestion = nextQuestion
	}
}

func askUser(questionText string) string {
	fmt.Println(questionText)
	fmt.Println("(Type 'back' to go to the previous question, 'start' for the first question, 'fork' for the last major fork, and 'exit' to end.)")
	fmt.Print("Your choice: ")

	var choice string
	fmt.Scanln(&choice)
	return strings.ToLower(choice)
}

func processResponse(response string, nextQuestion QuestionFunc) QuestionFunc {
	switch response {
	case "back":
		if previousQuestion != nil {
			return previousQuestion
		}
		fmt.Println("No previous question to go back to.")
	case "start":
		return startQuestion
	case "fork":
		if majorForkQuestion != nil {
			return majorForkQuestion
		}
		fmt.Println("No major fork to go back to.")
	case "exit":
		return nil
	default:
		return nextQuestion
	}
	return previousQuestion
}

func Q1() (string, AnswerMatrix) {
	question := `
Welcome to the decision tree bot!
Do you want to proceed?
1. Yes
2. No
`
	choices := AnswerMatrix{
		"1": Q2,
		"2": nil,
	}

	return question, choices
}

func Q2() (string, AnswerMatrix) {
	// Marking Q2 as a major fork.
	majorForkQuestion = Q2

	question := `
Have you used a decision tree bot before?
1. Yes
2. No
`
	choices := AnswerMatrix{
		"1": nil, // For this shortened version, both choices end the conversation
		"2": nil,
	}

	return question, choices
}

