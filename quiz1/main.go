package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	filename := flag.String("filename", "problems.csv", "filename of the csv")
	time := flag.Uint("time", 30, "time in seconds of how long to run the quiz")
	shouldWeShuffle := flag.Bool("isShuffle", false, "Do you want to shuffle the questions?")

	flag.Parse()

	dat, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("error reading the file %s", *filename)
	}
	fmt.Print(string(dat), time, shouldWeShuffle)

}
