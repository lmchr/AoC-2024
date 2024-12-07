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
		inputs = append(inputs, line)
	}
	return inputs
}

func removeFromList(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func Bool2int(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func indexOf[T comparable](element T, data []T) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func restoreField[T any](source [][]T) [][]T {
	arrayCopy := make([][]T, len(source))
	for i := range source {
		arrayCopy[i] = make([]T, len(source[i]))
		copy(arrayCopy[i], source[i])
	}
	return arrayCopy
}
