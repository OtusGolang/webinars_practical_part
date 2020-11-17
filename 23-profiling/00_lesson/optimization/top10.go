package optimization

import (
	"sort"
	"strings"
)

type WordCounter struct {
	word  string
	count int
}

func Top10(text string) []string {
	splittedText := strings.Fields(text)

	m := make(map[string]int)
	for _, word := range splittedText {
		m[word]++
	}

	wordCounter := make([]WordCounter, 0, len(m))
	for k, v := range m {
		wordCounter = append(wordCounter, WordCounter{
			word:  k,
			count: v,
		})
	}

	sort.Slice(wordCounter, func(i, j int) bool {
		return wordCounter[i].count > wordCounter[j].count
	})

	resLen := 10
	if resLen > len(wordCounter) {
		resLen = len(wordCounter)
	}
	res := make([]string, resLen)

	for i := range res {
		res[i] = wordCounter[i].word
	}

	return res
}
