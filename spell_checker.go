package spellchecker

import (
	"bufio"
	"errors"
	"os"
	"sort"
)

type SpellChecker struct {
	words []string
}

// Init() creates a new SpellChecker instance.
func Init(dictionaryPath string) (*SpellChecker, error) {
	words, err := loadDictionary(dictionaryPath)
	if err != nil {
		return nil, err
	}
	return &SpellChecker{words}, nil
}

// Check finds the closest words to the given word up to the specified number.
func (sc *SpellChecker) Check(word string, num int) []string {
	var wordDistances []struct {
		word     string
		distance int
	}

	for _, w := range sc.words {
		editDist := getEditDistance(word, w)
		wordDistances = append(wordDistances, struct {
			word     string
			distance int
		}{w, editDist})
	}

	sort.Slice(wordDistances, func(i, j int) bool {
		return wordDistances[i].distance < wordDistances[j].distance
	})

	var sortedWords []string
	for _, wd := range wordDistances {
		sortedWords = append(sortedWords, wd.word)
	}

	if len(sortedWords) <= num {
		return sortedWords
	}
	return sortedWords[:num]
}

func loadDictionary(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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

func getEditDistance(w, w1 string) int {
	w, w1, _ = checkLength(w, w1)
	m := createMatrix(w, w1)
	m = wagnerFisher(m)

	lastList := m[len(m)-1]
	editDistance := int(lastList[len(lastList)-1])
	return editDistance
}

func checkLength(w, w1 string) (string, string, error) {
	if len(w) <= 3 || len(w1) <= 3 {
		return w, w1, errors.New("too short word")
	}
	if len(w) < len(w1) {
		return w, w1, nil
	}
	return w1, w, nil
}

func createMatrix(word, word1 string) [][]byte {
	wordsMatrix := make([][]byte, len(word1)+2)
	for i := range wordsMatrix {
		wordsMatrix[i] = make([]byte, len(word)+2)
	}

	for y := 0; y < len(word)+2; y++ {
		if !(y == 0 || y == 1) {
			wordsMatrix[0][y] = word[y-2]
		}
	}

	for x := 0; x < len(word1)+2; x++ {
		if !(x == 0 || x == 1) {
			wordsMatrix[x][0] = word1[x-2]
		}
	}

	for y := 2; y < len(word)+2; y++ {
		wordsMatrix[1][y] = byte(y - 1)
	}

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

func wagnerFisher(m [][]byte) [][]byte {
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
