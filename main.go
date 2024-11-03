package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quote := getRandomQuote()
	fmt.Println("Випадкові речення:", quote)
}