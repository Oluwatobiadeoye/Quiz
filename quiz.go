package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	Question string
	Answer string
}

var fileName string

func init() {
	flag.StringVar(&fileName, "fileName", "problems.csv", "The name of the file where questions are loaded")
	flag.Parse()
}

func main() {

	var quizzes []Quiz
	var userScore int
	lines, err := ReadCsv(fileName)

	if err != nil {
		log.Fatalln("Could not open file at filePath" + fileName, err)
	}

	for _, line := range lines {
		quizzes = append(quizzes, Quiz{Question: line[0], Answer: line[1]})
	}

	fmt.Printf("You have %d questions to answer. Your time starts now!!", len(quizzes))
	fmt.Printf("You get 2 points for every correct answer and otherwise lose 1 point")

	for _, currentQuiz := range quizzes {
		fmt.Printf("What is the value of %s", currentQuiz.Question)

		if answer := ReadAnswer(); currentQuiz.Answer == answer {
			userScore += 2
		} else {
			userScore -= 1
		}
	}

	fmt.Printf("Your final score is %d out of possible %d", userScore, 2*len(quizzes))
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

func ReadAnswer() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Answer: ")
	text, _ := reader.ReadString('\n')
	return text
}