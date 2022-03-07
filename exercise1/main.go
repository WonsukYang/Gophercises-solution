package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question, answer string
}

func main() {
	filename := flag.String("csv", "problem.csv", "CSV file")
	file, err := os.Open(*filename)

	if err != nil {
		errf := fmt.Errorf("could not open file: %v", err)
		fmt.Println(errf.Error())
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		errf := fmt.Errorf("could not read file: %v", err)
		fmt.Println(errf.Error())
	}

	problems := []Problem{}
	for _, line := range lines {
		problems = append(problems, Problem{line[0], line[1]})
	}

	score := 0
	timer := time.NewTimer(time.Second * 2)

quizloop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChan := make(chan string)
		go func() {
			userAnswer := ""
			fmt.Scan(&userAnswer)
			answerChan <- strings.TrimSpace(userAnswer)
		}()

		select {
		case ans := <-answerChan:
			if ans == problem.answer {
				score++
			}
		case <-timer.C:
			fmt.Println("\nTimes up!")
			break quizloop
		}

	}
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}
