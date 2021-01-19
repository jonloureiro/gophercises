package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	answerChan := make(chan string)
	count := 0

	fileName := flag.String("f", "problems.csv", "file name")
	maxTime := flag.Int64("t", 30, "maximum time")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%ds time, ready? [enter]", *maxTime)
	bufio.NewReader(os.Stdin).ReadString('\n')
	timer := time.NewTimer(time.Duration(*maxTime) * time.Second)
	fmt.Println("----------")

	for _, v := range csvLines {
		fmt.Printf("%v: ", v[0])

		go func() {
			var input string
			_, err := fmt.Scanf("%v", &input)
			if err != nil {
				log.Println(err)
			}
			answerChan <- input
		}()

		select {
		case answer := <-answerChan:
			if strings.TrimSpace(answer) == strings.TrimSpace(v[1]) {
				count++
			}
		case <-timer.C:
			fmt.Println()
			break
		}
	}

	fmt.Println("----------")
	fmt.Println("Questions:", len(csvLines))
	fmt.Println("Correct:", count)
}
