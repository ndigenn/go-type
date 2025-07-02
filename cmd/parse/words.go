package parse

import (
	"io"
	"log"
	"os"
	"strings"
)

// map that filters structs of words based on size
var FilteredWords = map[string][]string {
	"small-10": {},
	"medium-25": {},
	"large-50": {},
}

func ParseWords() map[string][]string {
	// open file
	f, err := os.Open("cmd/parse/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read all characters
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// set bytes to a string and split the fields
	words := strings.Fields(string(b))

	total := len(words)
	for i := 0; i+10 <= total; i += 10 {
		FilteredWords["small-10"] = append(FilteredWords["small-10"], words[i:i+10]...)
	}

	for i := 0; i+25 <= total; i += 25 {
		FilteredWords["medium-25"] = append(FilteredWords["medium-25"], words[i:i+25]...)
	}

	for i := 0; i+50 <= total; i += 50 {
		FilteredWords["large-50"] = append(FilteredWords["large-50"], words[i:i+50]...)
	}

	return FilteredWords
}
