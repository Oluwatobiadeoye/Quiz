package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer string
}

var (
	fileName string
	quizDuration int
	answerChan  = make(chan string, 0)
)

func init() {
	flag.StringVar(&fileName, "file-location", "problems.csv", "The path the file where questions are loaded")
	flag.IntVar(&quizDuration, "duration", 30, "The duration of the quiz in seconds")
	flag.Parse()
}

func main() {

	var (
		quizzes []Quiz
		userScore int
		timeout <- chan time.Time
	)

	duration :=  time.Duration(quizDuration) * time.Second

	lines, err := ReadCsv(fileName)

	if err != nil {
		log.Fatalln("Could not open file at filePath" + fileName, err)
	}

	for _, line := range lines {
		quizzes = append(quizzes, Quiz{Question: line[0], Answer: line[1]})
	}

	fmt.Printf("You have %d questions to answer. Your time starts now!! \n\n", len(quizzes))
	fmt.Println("You get 2 points for every correct answer and otherwise lose 1 point")
	fmt.Printf("You have %s to answer all questions \n\n", duration)

	timeout = time.After(duration)

outer:
	for _, currentQuiz := range quizzes {
		fmt.Printf("What is the value of %s \n", currentQuiz.Question)
		go ReadAnswer()
			select {
				case <- timeout:
					fmt.Printf("\n\nYour time is up\n")
					break outer
				case answer := <- answerChan:
					if currentQuiz.Answer == answer {
						userScore += 2
					} else {
						userScore -= 1
					}
					fmt.Println()
					continue
			}
	}
	fmt.Printf("Your final score is %d out of possible %d \n", userScore, 2*len(quizzes))
}


func ReadCsv(filepath string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func ReadAnswer() {
	var answer string
	fmt.Print("Enter Answer: ")
	_, _ = fmt.Scanf("%s", &answer)
	answerChan <- answer
}