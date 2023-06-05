package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const liveSymbol = "\u2588"
const deadSymbol = " "

type World struct {
	grid       [][]int
	gridSize   int
	generation int
}

func (w *World) alive() (alive int) {
	for i := range w.grid {
		for j := range w.grid[i] {
			if w.grid[i][j] == 1 {
				alive++
			}
		}
	}
	return alive
}

func main() {
	gridSize, err := readInputs()
	if err != nil {
		panic(err)
	}
	w := makeWorld(gridSize)
	printWorld(w)

	exitChan := make(chan bool)
	defer close(exitChan)
	go func(w *World) {
		for {
			select {
			case _, canceled := <-exitChan:
				if canceled {
					return
				}
			default:
				time.Sleep(time.Millisecond * 100)
				tick(w)
				clearConsole()
				printWorld(w)
			}

		}
	}(w)
	listenExit(exitChan)
}

func tick(w *World) {
	var newGenGrid = make([][]int, w.gridSize)
	for i := range newGenGrid {
		newGenGrid[i] = make([]int, w.gridSize)
	}

	for i := range w.grid {
		for j := range w.grid[i] {
			liveCells := checkLiveNeighbours(w, i, j)
			if w.grid[i][j] == 1 {
				if liveCells == 2 || liveCells == 3 {
					newGenGrid[i][j] = 1
				}
				continue
			}
			if liveCells == 3 {
				newGenGrid[i][j] = 1
			}
		}
	}
	w.grid = newGenGrid
	w.generation++
}

func checkLiveNeighbours(w *World, cellI, cellJ int) (liveCells int) {
	if cellI < 0 || cellJ < 0 || cellI >= w.gridSize || cellJ >= w.gridSize {
		panic("Cell out of world")
	}

	for i := cellI - 1; i <= cellI+1; i++ {
		for j := cellJ - 1; j <= cellJ+1; j++ {
			if i == cellI && j == cellJ {
				continue
			}
			adjustedI := i
			adjustedJ := j
			if i == -1 {
				adjustedI = w.gridSize - 1
			}
			if i == w.gridSize {
				adjustedI = 0
			}
			if j == -1 {
				adjustedJ = w.gridSize - 1
			}
			if j == w.gridSize {
				adjustedJ = 0
			}
			if w.grid[adjustedI][adjustedJ] == 1 {
				liveCells++
			}
		}
	}
	return liveCells
}

func makeWorld(gridSize int) *World {
	grid := make([][]int, gridSize)
	for i := range grid {
		if grid[i] == nil {
			grid[i] = make([]int, gridSize)
		}
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}

	return &World{
		grid:     grid,
		gridSize: gridSize,
	}
}

func printWorld(w *World) {
	fmt.Printf("Generation #%v\nAlive: %v\n", w.generation, w.alive())

	printBoxBorder(w.gridSize, true)
	for i := range w.grid {
		for j := range w.grid[i] {
			if j == 0 {
				fmt.Printf("%v", "\u2503")
			}
			cell := deadSymbol
			if w.grid[i][j] == 1 {
				cell = liveSymbol
			}
			fmt.Printf("%v", cell)

			if j == w.gridSize-1 {
				fmt.Printf("%v", "\u2503")
			}
		}
		fmt.Printf("\n")
	}
	printBoxBorder(w.gridSize, false)
}

func printBoxBorder(size int, top bool) {
	for i := 0; i < size+2; i++ {
		if i == 0 {
			if top {
				fmt.Printf("%v", "\u250f")
			} else {
				fmt.Printf("%v", "\u2517")
			}
			continue
		}
		if i == size+1 {
			if top {
				fmt.Printf("%v", "\u2513")
			} else {
				fmt.Printf("%v", "\u251b")
			}
			continue
		}
		fmt.Printf("%v", "\u2501")
	}
	fmt.Printf("\n")
}

func readInputs() (gridSize int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Grid size > ")
	for scanner.Scan() {
		nStr := strings.TrimSpace(scanner.Text())
		gridSize, err := strconv.ParseInt(nStr, 10, 32)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return int(gridSize), nil

	}

	if scanner.Err() != nil {
		return 0, err
	}
	panic("Can't read inputs")
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func listenExit(exitChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exitChan <- true
		return
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
