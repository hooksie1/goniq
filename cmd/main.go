package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/hooksie1/goniq"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "file to use")
}

func main() {
	flag.Parse()

	file, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	list := goniq.Uniq(file)

	for _, v := range list {
		fmt.Printf("%v \n", v)
	}

}
