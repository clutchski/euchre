package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	d := NewDeckShuffled()
	for _, card := range d {
		fmt.Println(card)
	}

}
