package main

import (
	"fmt"
	"strings"
)

func mapMaker (str string) map[string]int {
	wordMap := make(map[string]int)
	strArr := strings.Split(str, "")
	for i := 0; i < len(strArr); i++ {
		wordMap[strArr[i]] ++
	}
	return wordMap
}
func Anagrams(word string, words []string) []string {
	// your code
	anagrams := []string {}
	word1 := mapMaker(word)
	for i := 0; i < len(words); i++ {
		matches := 0
		word2 := mapMaker(words[i])
		if len(word1) == len(word2) && len(word) == len(words[i]) {
			for key := range word1 {
				if strings.Contains(words[i], key) {
					matches++
				}
				if matches == len(word1) {
					anagrams = append(anagrams, words[i])
				}
			}
		}
	}
	if len(anagrams) == 0 {
		return nil
	}
	return anagrams
}

func main() {
	fmt.Println(Anagrams("racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"}))
}