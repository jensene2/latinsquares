package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	// Check for the right amount of arguments.
	if len(os.Args) != 3 {
		log.Fatal("Not enough arguments. LatinSquare size and output filename are required.")
	}

	// Parse the size into an integer.
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Open up the output file.
	output, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	// Defer closing of the file once everything has finished.
	defer output.Close()

	// Wrap the file in a buffered writer.
	writer := bufio.NewWriter(output)

	// Start getting the latin squares of the given size.
	// Results is a channel of all of the valid latin squares.
	results := start(i)

	// Loop through all of the results as they are received.
	// Count them and write them to a file.
	count := 0
	for result := range results {
		count++

		_, err = writer.WriteString(result.String())
		if err != nil {
			log.Fatal(err)
		}

		writer.Flush()
	}

	// Just print the final count.
	fmt.Println(count)
}

func start(size int) <-chan LatinSquare {
	// Create a new empty latin square.
	square := NewReducedLatinSquare(size)

	// Make the results channel what will be returned.
	results := make(chan LatinSquare)

	// Create a wait group that'll be passed down.
	wg := new(sync.WaitGroup)

	// Start a goroutine to start the permutation process and close
	//   the results channel once it has finished.
	go func() {
		next(square, results, wg)
		wg.Wait()
		close(results)
	}()

	return results
}

func next(square LatinSquare, results chan<- LatinSquare, wg *sync.WaitGroup) {
	// If the square is finished, send it to the results and return.
	if square.isFinished() {
		results <- square
		return
	}

	// Get the first unset coordinates.
	x, y := square.getFirstUnsetCoordinates()

	// Get the possibilities for the current coordinates.
	possibilities := square.getPossibilities(x, y)

	// Create a wait group so this function doesn't return until it's
	//   children have returned.
	//wg := new(sync.WaitGroup)
	wg.Add(len(possibilities))

	// Create a goroutine for each child.
	for _, value := range possibilities {
		// Make a copy and set the value.
		newSquare := square.copy()
		newSquare.set(x, y, value)

		// Create the goroutine. Not only does this call the function
		//   recursively, but it will also signal that it is done to the
		//   wait group.
		go func(square LatinSquare) {
			next(square, results, wg)
			wg.Done()
		}(newSquare)
	}
}
