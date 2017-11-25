package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func printResults(numCorrect int, numIncorrect int) {
	fmt.Printf("Number correct: %v, number incorrect: %v\n", numCorrect, numIncorrect)
}

func main() {
	finished := make(chan bool)
	file, err := os.Open("./problems.csv")

	if err != nil {
		panic(err)
	}

	numCorrect := 0
	numIncorrect := 0

	r := csv.NewReader(bufio.NewReader(file))

	go func() {
		for {
			record, err := r.Read()
			if err == io.EOF {
				finished <- true
			}
			if err != nil {
				log.Fatal(err)
			}

			problem := record[0]
			solution := record[1]

			fmt.Printf("%v?\n", problem)
			var response string
			fmt.Scanln(&response)

			if response == solution {
				fmt.Println("Correct!")
				numCorrect++
			} else {
				fmt.Println("Incorrect :(")
				numIncorrect++
			}
		}
	}()

	timeLimit := time.After(30 * time.Second)

	for {
		select {
		case <-timeLimit:
			fmt.Println("TIMES UP YER DONE")
			printResults(numCorrect, numIncorrect)
			os.Exit(0)
		case <-finished:
			fmt.Println("Congrats you finished!")
			printResults(numCorrect, numIncorrect)
			os.Exit(0)
		}
	}
}
