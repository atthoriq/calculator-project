package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/atthoriq/calculator-project/calculator"
)

func main() {
	fmt.Println("Welcome to The Calculator!")
	defer fmt.Println("Good bye!")

	// initialize handler and dependencies
	calc := calculator.InitNewCalculator()
	handler := InitCalculatorHandler(calc)
	if handler == nil {
		log.Fatal("fail initializing handler")
	}

	// run the scanner
	inputScanner(handler)
}

func inputScanner(handler *calculatorHandler) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			result, err := handler.Handle(text)
			if err != nil {
				log.Fatal("error occurred: ", err)
				break
			}
			if len(result) == 0 {
				break
			}

			fmt.Println(result)
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		log.Fatal("Error: ", scanner.Err())
	}
}
