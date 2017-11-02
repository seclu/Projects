package data_structures

import (
	"regexp"
	"strings"
)

type Search struct {
	Text string
}

func (s *Search) Search(query string) int {
	dict := s.createInverseIndex(s.Text)

	if count, ok := dict[query]; ok {
		return count
	}

	return 0
}

func (s *Search) createInverseIndex(text string) map[string]int {
	var dict = make(map[string]int)

	lowerText := strings.ToLower(text)

	alpha := regexp.MustCompile("[^a-zA-Z\\s+]+")
	words := regexp.MustCompile("\\s+")

	onlyWords := alpha.ReplaceAllString(lowerText, " ")

	for _, word := range words.Split(onlyWords, -1) {
		if len(word) < 1 {
			continue
		}

		dict[word]++
	}

	return dict
}