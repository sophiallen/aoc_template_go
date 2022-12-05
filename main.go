package main

import (
    "bufio"
    "fmt"
    "os"
	"aoc_template_go/template"
)

type Solvable interface {
	Solve(scanner *bufio.Scanner)
	GetDataPath() string
}

var solutions = map[int]Solvable{
	0: template.Solution{},
}

func main() {
	const day = 1
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
