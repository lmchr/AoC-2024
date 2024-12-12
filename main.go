package main

import (
	"aoc-2024/Days"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	var day int
	var cpuProfile string
	flag.IntVar(&day, "day", -1, "Provide the day to execute [1..24]")
	flag.StringVar(&cpuProfile, "cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	log.Printf("AoC-2024")
	start := time.Now()
	switch day {
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
	case 8:
		Days.Day8()
	case 9:
		Days.Day9()
	case 10:
		Days.Day10()
	default:
		panic(fmt.Sprintf("Invalid day or not implemented yet: %d. Execute with e.g. '-day 1'", day))
	}
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
