package Days

import (
	"bufio"
	"log"
	"os"
)

func readFile(filename string) []string {
	var inputs []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			inputs = append(inputs, line)
		}
	}
	return inputs
}

func removeFromList(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
