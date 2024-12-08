package main

import (
	"aoc-2024/Days"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var intFlag int
	flag.IntVar(&intFlag, "day", -1, "Provide the day to execute [1..24]")
	flag.Parse()
	log.Printf("AoC-2024")
	start := time.Now()
	switch intFlag {
	case 1:
		Days.Day1()
	case 2:
		Days.Day2()
	case 3:
		Days.Day3()
	case 4:
		Days.Day4()
	case 5:
		Days.Day5()
	case 6:
		Days.Day6()
	case 7:
		Days.Day7()
	default:
		panic(fmt.Sprintf("Invalid day or not implemented yet: %d. Execute with e.g. '-day 1'", intFlag))
	}
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
