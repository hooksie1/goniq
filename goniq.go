package goniq

import (
	"bytes"
	"io"
)

func toBytes(r io.Reader) []byte {
	data := new(bytes.Buffer)

	data.ReadFrom(r)
	return data.Bytes()
}

// uniq takes a list of items only shows the unique items.
func Uniq(r io.Reader) io.Reader {
	bSlice := toBytes(r)

	lstmap := map[byte]bool{}
	result := []byte{}

	for v := range bSlice {
		if lstmap[bSlice[v]] == true {

		} else {
			lstmap[bSlice[v]] = true
			result = append(result, bSlice[v])
		}
	}
	return bytes.NewReader(result)
}
