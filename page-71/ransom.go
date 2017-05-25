// Example: A ransom note can be formed by cutting words
// out of a magazine to form a new sentence. How would you
// figure out if a ransom note (represented as a string)
// can be formed from a given magazine (string)?

package main

import (
	"fmt"
	"strings"
)

func IsPossibleRansom(ransom, magazine string) bool {
	// turn ransom into a hash table of words to count
	ransomWordCount := make(map[string]int)
	for _, word := range strings.Fields(ransom) {
		ransomWordCount[word] += 1
	}
	// turn magazine into a hash table of words to count
	magazineWordCount := make(map[string]int)
	for _, word := range strings.Fields(magazine) {
		magazineWordCount[word] += 1
	}
    // check if enough words exist in magazine to form ransom
	for word, count := range ransomWordCount {
		if magazineWordCount[word] < count {
			return false
		}
	}
	return true
}

func main() {
	ransom := "hello world this"
	magazine := "hello this is a world"
	fmt.Println(IsPossibleRansom(ransom, magazine))
	notransom := "hello cats"
	fmt.Println(IsPossibleRansom(notransom, magazine))
	notransom = "hello thi"
	fmt.Println(IsPossibleRansom(notransom, magazine))
}
