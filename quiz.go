package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./problems.csv")

	if err != nil {
		panic(err)
	}

	numCorrect := 0
	numIncorrect := 0

	r := csv.NewReader(bufio.NewReader(file))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
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

	fmt.Printf(" Number correct: %v, number incorrect: %v\n", numCorrect, numIncorrect)
}
