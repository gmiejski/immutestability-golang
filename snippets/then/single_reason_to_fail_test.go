package then

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatchesAllSentencesWithExactWordsOccurrence(t *testing.T) {
	// given
	searchFacade := newInMemorySearch([]string{"dogs cats", "no animal", "cats are smart"})

	// when
	results := searchFacade.Search("cat")

	// then
	assert.ElementsMatch(t, []string{"dogs cats", "cats are smart"}, extractSentences(results))
}

func TestSearchReturnsMatchedWords(t *testing.T) {
	// given
	searchFacade := newInMemorySearch([]string{"dogs cats"})

	// when
	results := searchFacade.Search("cat")

	// then
	assert.Equal(t, 1, len(results.matches))
	assert.Equal(t, []string{"cats"}, results.matches[0].wordsMatched)
}

func TestSearchResults(t *testing.T) { // 1. it's hard to find a name for the test, as it tests 2 behaviours at once
	// given
	searchFacade := newInMemorySearch([]string{"dogs cats", "no animal", "cats are smart"})

	// when
	results := searchFacade.Search("cat")

	// then
	assert.ElementsMatch(t, []SearchResult{
		{sentence: "dogs cats", wordsMatched: []string{"cats"}},
		{sentence: "cats are smart", wordsMatched: []string{"cats"}},
	}, results.matches)
	// 2. test will fail if we have bug in either of the features - searching and returning the matched words
	// 3. test will fail if we change matched keywords logic - even though returning sentences still works
}

func extractSentences(results SearchResults) []string {
	var result []string
	for _, x := range results.matches {
		result = append(result, x.sentence)
	}
	return result
}

type SearchResults struct {
	matches []SearchResult
}

type SearchResult struct {
	sentence     string
	wordsMatched []string
}

type SearchFacade interface {
	Search(word string) SearchResults
}

func newInMemorySearch(sentences []string) SearchFacade {
	return &searchFacade{sentences: sentences}
}

type searchFacade struct {
	sentences []string
}

func (s *searchFacade) Search(word string) SearchResults {
	// TODO
	return SearchResults{}
}
