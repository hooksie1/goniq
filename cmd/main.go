package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitlab.com/hooksie1/goniq"
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
	list := goniq.Uniq(output)
	for _, v := range list {
		fmt.Printf("%v \n", v)
	}

}
