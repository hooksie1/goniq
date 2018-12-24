package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines takes a file name and returns the contents as a byte slice.
func readLines(p string) []string {
	file, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// uniq takes a list of items only shows the unique items.
func uniq(s []string) []string {
	lstmap := map[string]bool{}
	result := []string{}

	for v := range s {
		if lstmap[s[v]] == true {

		} else {
			lstmap[s[v]] = true
			result = append(result, s[v])
		}
	}
	return result
}

func main() {
	if len(os.Args) > 1 {
		file := os.Args[1]
		lines := readLines(file)
		list := uniq(lines)

		for _, v := range list {
			fmt.Printf("%v \n", v)
		}
	} else {
		var text string
		var bs []string
		var output []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text = scanner.Text()
			bs = strings.Split(text, " ")
			for _, v := range bs {
				output = append(output, v)
			}
		}
		list := uniq(output)
		for _, v := range list {
			fmt.Printf("%v \n", v)
		}

	}

}
