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
		word2 := mapMaker(words[i])
		if len(word1) == len(word2) {
			for key := range word1 {
				_, inWord2 := word2[key]
				if !inWord2 {
					fmt.Println("key", i)
				}
				if word1[key] != word2[key] {
					break
				}
			}
			fmt.Println(words[i], i)
			anagrams = append(anagrams, words[i])
		}
	}
	return anagrams
}

func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"}))
}