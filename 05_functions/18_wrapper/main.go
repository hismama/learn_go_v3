package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("./05_functions/18_wrapper/poem.txt")
	if err != nil {
		log.Fatalf("error in main in readFile: %v", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}
