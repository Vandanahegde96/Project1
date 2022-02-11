package service

import (
	"strings"
	"sort"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}


func CountStringsInText(input string)  []WordCount{
	m := make(map[string]int)
    var Words []WordCount

	// convert the input string to array of words
	arr := strings.Fields(input)
	for _, a := range arr {
		m[a]++
	}

	for k,v := range m {
		Words = append(Words, WordCount{Word: k, Count: v})
	}

	 // Sort the Array of struct by using go's built in library soryt
    sort.Slice(Words, func(i,j int) bool {
		return Words[i].Count > Words[j].Count
	})

	if len(Words) < 10{
	   return Words
	}
	return Words[0:10]

}