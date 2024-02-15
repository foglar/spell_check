package main

import (
	"fmt"
	"github.com/foglar/spell_check/spell"
	"log"
)

func main() {
	sc, err := spellchecker.NewSpellChecker("./words.txt")
	if err != nil {
		log.Fatalf("Error creating SpellChecker: %v", err)
	}

	word := "exprezzion"
	closestWords := sc.Check(word, 10)

	fmt.Printf("Closest words to '%s': %v\n", word, closestWords)
}
