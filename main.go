package main

import (
	"aoc_template_go/template"
	"bufio"
	"fmt"
	"os"
)

type Solvable interface {
	Solve(scanner *bufio.Scanner)
	GetDataPath() string
}

var solutions = map[int]Solvable{
	0: template.Solution{},
	// add each day following the pattern above
}

func main() {
	// change day value to switch the solution that runs
	const day = 0
	solution := solutions[day]
	fileScanner, reader := openFile(solution.GetDataPath())
	solution.Solve(fileScanner)
	reader.Close()
}

func openFile(file string) (*bufio.Scanner, *os.File) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readFile
}
