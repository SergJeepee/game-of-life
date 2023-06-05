package game

import (
	"fmt"
	"math/rand"
)

const liveSymbol = "\u2588"
const deadSymbol = " "

type World struct {
	grid       [][]int
	gridHeight int
	gridWidth  int
	generation int
}

func (w *World) Print() {
	clearConsole()
	fmt.Printf("Generation #%v\nAlive: %v\n", w.generation, w.Alive())

	printBoxBorder(w.gridWidth, true)
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

			if j == w.gridWidth-1 {
				fmt.Printf("%v", "\u2503")
			}
		}
		fmt.Printf("\n")
	}
	printBoxBorder(w.gridWidth, false)
}

func (w *World) Alive() (alive int) {
	for i := range w.grid {
		for j := range w.grid[i] {
			if w.grid[i][j] == 1 {
				alive++
			}
		}
	}
	return alive
}

func (w *World) Tick() {
	var newGenGrid = make([][]int, w.gridHeight)
	for i := range newGenGrid {
		newGenGrid[i] = make([]int, w.gridWidth)
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

func MakeWorld(gridSize int) *World {
	//making more square looking world
	gridHeight := gridSize
	gridWidth := gridHeight * 3

	grid := make([][]int, gridHeight)
	for i := range grid {
		if grid[i] == nil {
			grid[i] = make([]int, gridWidth)
		}
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}

	return &World{
		grid:       grid,
		gridWidth:  gridWidth,
		gridHeight: gridSize,
	}
}

func checkLiveNeighbours(w *World, cellI, cellJ int) (liveCells int) {
	if cellI < 0 || cellJ < 0 || cellI >= w.gridHeight || cellJ >= w.gridWidth {
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
				adjustedI = w.gridHeight - 1
			}
			if i == w.gridHeight {
				adjustedI = 0
			}
			if j == -1 {
				adjustedJ = w.gridWidth - 1
			}
			if j == w.gridWidth {
				adjustedJ = 0
			}
			if w.grid[adjustedI][adjustedJ] == 1 {
				liveCells++
			}
		}
	}
	return liveCells
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

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}
