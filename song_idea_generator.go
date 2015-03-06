package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Use the time to initialize the default Source to a
	// deterministic state
	rand.Seed(time.Now().UTC().UnixNano())

	// Create the words
	themes := []string{
		"love",
		"heartbreak",
		"death",
	}

	keys := []string{
		"E",
		"D",
		"A minor",
	}

	meter := []string{
		"4/4",
		"3/4",
		"mixed",
	}

	inst := []string{
		"full band",
		"guitar and voice",
		"electronic",
		"no",
	}

	// Generate an idea
	fmt.Println("GENIUS IDEA:")
	fmt.Println("A song about", themes[randInt(len(themes))])
	fmt.Println("In the key of", keys[randInt(len(keys))])
	fmt.Println("With", meter[randInt(len(meter))], "meter")
	fmt.Println("And", inst[randInt(len(keys))], "instrumentation")
}

// Return a random number
func randInt(size int) int {
	return rand.Intn(size)
}
