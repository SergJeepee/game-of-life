package main

import (
	"bufio"
	"fmt"
	"github.com/SergJeepee/gameoflife/game"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var gameInProgress = false

func main() {
	restartChan := make(chan bool)
	for {
		go listenControls(restartChan)
		runGame(restartChan)
	}
}

func runGame(restartChan chan bool) {
	size, err := readInputs()
	if err != nil {
		panic(err)
	}
	w := game.MakeWorld(size)
	w.Print()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(w *game.World) {
		defer wg.Done()
		for {
			select {
			case _, canceled := <-restartChan:
				if canceled {
					return
				}
			default:
				time.Sleep(time.Millisecond * 50)
				w.Tick()
				w.Print()
			}
		}
	}(w)
	gameInProgress = true
	wg.Wait()
	gameInProgress = false
	return
}

func readInputs() (worldInput int, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Available presets: [%v]\n", strings.Join(game.AvailablePresets(), ", "))
	fmt.Print("Grid size of random world or preset name > ")
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(input, "exit") {
			os.Exit(0)
		}
		preset := game.FindPreset(input)
		if preset != nil {
			return preset.Id, nil
		}
		worldInput, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if worldInput > 50 {
			fmt.Print("So big expectations, huh? Make it smaller > ")
			continue
		}
		return int(worldInput), nil

	}

	if scanner.Err() != nil {
		return 0, err
	}
	panic("Can't read inputs")
}

func listenControls(restartChan chan bool) {
	for !gameInProgress {
		//do nothing
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(input, "exit") {
			os.Exit(0)
		}
		restartChan <- true
		return
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
