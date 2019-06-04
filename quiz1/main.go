package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filename := flag.String("filename", "problems.csv", "filename of the csv")
	//time := flag.Uint("time", 30, "time in seconds of how long to run the quiz")
	//shouldWeShuffle := flag.Bool("isShuffle", false, "Do you want to shuffle the questions?")

	flag.Parse()

	dat, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("error reading the file %s\n", *filename)
	}
	qCount := 0
	r := csv.NewReader(strings.NewReader(string(dat)))
	questions := make([][]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		qCount = qCount + 1
		questions = append(questions, record)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Quiz")
	fmt.Println("---------------------")

	corrects := 0
	wrongs := 0

	for _, ques := range questions {
		fmt.Print(ques[0], " = ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare(ques[1], strings.TrimSpace(text)) == 0 {
			corrects = corrects + 1
		} else {
			wrongs = wrongs + 1
		}
	}

	fmt.Println(" ---------------------- DONE!")
	fmt.Println("The results are below:")
	fmt.Printf("You answered %d out of %d questions correctly!\n\n", corrects, qCount)
}
