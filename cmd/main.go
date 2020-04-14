package main

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/hooksie1/goniq"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileList := goniq.Uniq(file)

	list := goniq.Uniq(strings.NewReader("test,test2,test3,test2"))
	fmt.Println(list)
	fmt.Println(fileList)

}
