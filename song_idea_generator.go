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
	fmt.Println("A song about", sample(themes))
	fmt.Println("In the key of", sample(keys))
	fmt.Println("With", sample(meter), "meter")
	fmt.Println("And", sample(inst), "instrumentation")
}

// Return a random array value
func sample(array []string) string {
	return array[rand.Intn(len(array))]
}
