//Part 1
// Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

// The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

// The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

// Flags package, csvs, strings and files
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"flag"
)

type Question struct {
	Problem string
	Answer  string
}

func QuestionPrompt(question string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, question+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)

}

func Timer(time int){
	if time == flag.NArg(){

	}
	flag.Int("time",time, )
	timer2 := time.NewTimer(time.Second)
}


func main() {
	// open file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var count int = 0
	q := &Question{
		Problem: "Your problem here",
		Answer:  "Your answer here",
	}
	for i := 0; i < len(data); i++ {
		rec := q
		rec.Problem = data[i][0]
		rec.Answer = data[i][1]

		answer := QuestionPrompt(rec.Problem)
		if rec.Answer == answer {
			count++
		}
	}

	// print the array
	fmt.Println("Number of correct answers", count)
	fmt.Println("Number of wrong answers", len(data)-count)
}
