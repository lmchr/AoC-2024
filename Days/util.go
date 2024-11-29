package Days

import (
	"bufio"
	"log"
	"os"
)

func readFile(filename string) []string {
	log.Println("Reading ", filename)
	var inputs []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}
	return inputs
}
