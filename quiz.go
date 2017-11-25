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

func main() {
	finished := make(chan struct{})
	file, err := os.Open("./problems.csv")

	if err != nil {
		panic(err)
	}

	numCorrect := 0
	numIncorrect := 0

	r := csv.NewReader(bufio.NewReader(file))

	fmt.Println("Press enter when you're ready!")
	fmt.Scanln()

	go func(r *csv.Reader, correct, incorrect *int) {
		for {
			record, err := r.Read()
			if err == io.EOF {
				finished <- struct{}{}
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
				*correct++
			} else {
				fmt.Println("Incorrect :(")
				*incorrect++
			}
		}
	}(r, &numCorrect, &numIncorrect)

	timeLimit := time.After(30 * time.Second)

	select {
	case <-timeLimit:
		fmt.Println("TIMES UP YER DONE")
	case <-finished:
		fmt.Println("Congrats you finished!")
	}

	fmt.Printf("Number correct: %v, number incorrect: %v\n", numCorrect, numIncorrect)
}
