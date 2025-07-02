package parse

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// maybe eventually add a topics selection so the user can choose a type of quote
// capital var names means public, lowercase is private
type Quote struct {
	Quote	string	 `json:"quote"`
	Author	string	 `json:"author"`
}

func ParseJSON() []Quote {
	f, err := os.Open("cmd/parse/quotes.json")
	if err != nil {
		// should have some hardcoded backups if can't open quotes file
		log.Fatal(err)
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var q []Quote

	if err :=json.Unmarshal(b, &q); err != nil {
		log.Fatal(err)
	}

	return q
}

