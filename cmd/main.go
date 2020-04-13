package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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
	list := goniq.uniq(output)
	for _, v := range list {
		fmt.Printf("%v \n", v)
	}

}
