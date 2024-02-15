package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

// TODO: Add ability to pick a dictionary you want to use
// TODO: CMD arguments

func main() {
	closest_words := SpellCheck("spade", 5)
	fmt.Print(closest_words)
	//closest_words := SpellCheck("ad", 10)
	//fmt.Println(closest_words)
}

func SpellCheck(word string, num int) []string {
	var words []string
	var err error

	words, err = LoadDictionary("./words.txt")
	if err != nil {
		fmt.Println("Error reading dictionary")
	}

	// Create a slice to store word-distance pairs
	var wordDistances []struct {
		word     string
		distance int
	}

	// Calculate edit distance for each word and store it along with the word
	for _, w := range words {
		editDist := GetEditDistance(word, w)
		wordDistances = append(wordDistances, struct {
			word     string
			distance int
		}{w, editDist})
	}

	// Sort the word-distance pairs based on distance
	sort.Slice(wordDistances, func(i, j int) bool {
		return wordDistances[i].distance < wordDistances[j].distance
	})

	// Extract the sorted words into a separate slice
	var sortedWords []string
	for _, wd := range wordDistances {
		sortedWords = append(sortedWords, wd.word)
	}

	// Return the sorted words up to the specified number
	if len(sortedWords) <= num {
		return sortedWords
	}
	return sortedWords[:num]
}

func GetEditDistance(w string, w1 string) int {
	w, w1, _ = check_length(w, w1)
	m := create_matrix(w, w1)
	m = WagnerFisher(m)

	last_list := m[len(m)-1]
	edit_distance := last_list[len(last_list)-1]
	return int(edit_distance)
}

func LoadDictionary(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func check_length(w, w1 string) (shortWord, longWord string, err error) {
	if len(w) <= 3 || len(w1) <= 3 {
		return w, w1, errors.New("Too short word")
	}
	if len(w) < len(w1) {
		return w, w1, nil
	}
	return w1, w, nil
}

func create_matrix(word string, word1 string) [][]byte {
	// Create a matrix of x = len(word); y = len(word1)
	wordsMatrix := make([][]byte, len(word1)+2)
	for i := range wordsMatrix {
		wordsMatrix[i] = make([]byte, len(word)+2)
	}

	for y := 0; y < len(word)+2; y++ {
		if !(y == 0 || y == 1) {
			wordsMatrix[0][y] = word[y-2]
		}

	}
	// Add a word to a column
	for x := 0; x < len(word1)+2; x++ {
		if !(x == 0 || x == 1) {
			wordsMatrix[x][0] = word1[x-2]
		}
	}

	// Add first number row
	for y := 2; y < len(word)+2; y++ {
		wordsMatrix[1][y] = byte(y - 1)
	}

	// Add first number column
	for x := 2; x < len(word1)+2; x++ {
		wordsMatrix[x][1] = byte(x - 1)
	}

	return wordsMatrix
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func WagnerFisher(m [][]byte) [][]byte {
	for i := 2; i < len(m); i++ {
		for j := 2; j < len(m[0]); j++ {
			if m[i][0] != m[0][j] {
				m[i][j] = byte(min(int(m[i][j-1])+1, int(m[i-1][j])+1, int(m[i-1][j-1])+1))
			} else {
				m[i][j] = m[i-1][j-1]
			}
		}
	}
	return m
}
