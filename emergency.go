package main

import (
	"fmt"
	"os"
)

func main() {
	// startNumDiv()
	startFindWordsInGrid()
}

type grid [][]rune

func startFindWordsInGrid() {
	testCase := 0
	fmt.Scanln(&testCase)
	if testCase == 0 {
		os.Exit(0)
	}
	gridColl := make([]grid, 0)
	wordColl := make([]string, 0)
	for index := 1; index <= testCase; index++ {
		width, length := 0, 0
		tempWordGrid, tempWord := "", ""
		fmt.Scanln(&width)
		fmt.Scanln(&length)
		tempGrid := make(grid, width)
		for index := range tempGrid {
			tempGrid[index] = make([]rune, length)
		}
		for indexGrid := 0; indexGrid < width; indexGrid++ {
			fmt.Scanln(&tempWordGrid)
			for indexWord, value := range tempWordGrid {
				tempGrid[indexGrid][indexWord] = value
			}
		}
		gridColl = append(gridColl, tempGrid)
		fmt.Scanln(&tempWord)
		wordColl = append(wordColl, tempWord)
	}
	fmt.Println("Grid Collection : ", gridColl)
	fmt.Println("Word Collection : ", wordColl)
}

func startNumDiv() {
	testCase := 0
	fmt.Scanln(&testCase)
	if testCase == 0 {
		os.Exit(0)
	}
	result := make([]int, 0)
	for index := 1; index <= testCase; index++ {
		start, finish, divisor := 0, 0, 0
		fmt.Scanln(&start)
		fmt.Scanln(&finish)
		fmt.Scanln(&divisor)
		result = append(result, numDiv(start, finish, divisor))
	}
	for index, value := range result {
		if index == len(result)-1 {
			fmt.Printf("Case %d: %d", index+1, value)
		} else {
			fmt.Printf("Case %d: %d\n", index+1, value)
		}
	}
}
