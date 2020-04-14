package goniq

import (
	"bytes"
	"io"
)

// toBytes takes an io.Reader and converts it to a byte slice.
func toBytes(r io.Reader) []byte {
	data := new(bytes.Buffer)

	data.ReadFrom(r)
	return data.Bytes()
}

// splitWords takes a byte slice and returns a slice of byte slices
// separated by spaces, to be converted to a slice of strings.
func splitWords(b []byte) [][]byte {

	split := bytes.Split(b, []byte(","))

	return split

}

// Uniq takes an io.Reader containing a list of space separated items,
// and returns a slice of strings containing only the unique items.
func Uniq(r io.Reader) []string {
	bSlice := toBytes(r)

	lstmap := map[string]bool{}
	result := []string{}

	words := splitWords(bSlice)

	for _, v := range words {
		word := string(v)
		if lstmap[word] == true {

		} else {
			lstmap[word] = true
			result = append(result, string(v))
		}
	}

	return result

}
