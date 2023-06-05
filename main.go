package main

import (
	"bufio"
	"fmt"
	"github.com/SergJeepee/gameoflife/game"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	size, err := ReadInputs()
	if err != nil {
		panic(err)
	}
	w := game.MakeWorld(size)
	w.Print()

	exitChan := make(chan bool)
	defer close(exitChan)
	go func(w *game.World) {
		for {
			select {
			case _, canceled := <-exitChan:
				if canceled {
					return
				}
			default:
				time.Sleep(time.Millisecond * 100)
				w.Tick()
				w.Print()
			}

		}
	}(w)
	listenExit(exitChan)
}

func ReadInputs() (size int, err error) {
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
