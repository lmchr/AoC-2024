package main

import (
	"aoc-2024/Days"
	"log"
	"time"
)

func main() {
	log.Printf("AoC-2024")
	start := time.Now()
	Days.Day1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
	start = time.Now()
	Days.Day2()
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
	start = time.Now()
	Days.Day3()
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
	start = time.Now()
	Days.Day4()
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
	start = time.Now()
	Days.Day5()
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
	start = time.Now()
	Days.Day6()
	elapsed = time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
