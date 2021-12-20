package utils

import (
	"bufio"
	"log"
	"os"
)

type Input struct {
	I int
	S string
	B bool
}

func ReadFileByLines(path string, il *[]Input, proc func(line string) Input) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*il = append(*il, proc(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
