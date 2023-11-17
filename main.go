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

func QuestionPrompt() string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)

}

func Timer() int{

	timeptr := flag.Int("time", 30, "an int")
	flag.Parse()
	return *timeptr
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
	fmt.Println("Press Enter to start the timer...")
	QuestionPrompt() // Wait for Enter key press to start the timer

	timeValue := Timer()
	timer := time.NewTimer(time.Duration(timeValue) * time.Second)


	problemloop:
	
		for i := 0; i < len(data); i++ {
			fmt.Printf("Problem %d: %s = ",i+1, data[i][0])
			answerCh := make(chan string)
			go func ()  {
				answer := QuestionPrompt()
				answerCh <- answer
			}()

			select {
				case <-timer.C:
					fmt.Printf("\nScore: %d out of %d", count, len(data))
					break problemloop
				case answer := <- answerCh:
					rec := q
					rec.Problem = data[i][0]
					rec.Answer = data[i][1]
					if rec.Answer == answer {
						count++
					}

			}
		}



	// print the array
	fmt.Printf("\nScore: %d out of %d", count, len(data))
}
