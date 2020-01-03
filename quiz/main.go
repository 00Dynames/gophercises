package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	csvFlag := flag.String("q", "problems.csv", "A csv file containing questions and answers")
	timerFlag := flag.Int("t", 30, "Time limit for quiz completion in seconds")
	randFlag := flag.Bool("r", false, "Randomise question order")
	flag.Parse()

	csvFile, _ := os.Open(*csvFlag)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	lines, _ := reader.ReadAll()
	problems := parse(lines)

	questions := []string{}
	correct := 0
	timer := time.NewTimer(time.Duration(*timerFlag) * time.Second)

	if *randFlag {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

problemloop:
	for _, p := range problems {

		fmt.Printf(p.q)
		answers := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", answer)
			answers <- strings.Trim(answer, " ")
		}()

		select {
		case <-timer.C:
			break problemloop
		case ans := <-answers:
			if ans == p.a {
				correct++
			}
		}

	}

	fmt.Printf("\nYou got %d correct out of %d\n", correct, len(questions))
}

func parse(lines [][]string) []problem {

	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return result
}

type problem struct {
	q string
	a string
}
