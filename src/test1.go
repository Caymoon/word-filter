package main

import (
	"filter"
	"fmt"
	_ "io/ioutil"
	"time"
	_ "unicode/utf8"
)

func main() {

	var str []rune
	var dicFiles []string

	dicFile1 := "dic.txt"

	dicFiles = append(dicFiles, dicFile1)

	filter.LoadDicFiles(dicFiles)
	time.Sleep(1000 * time.Millisecond) //asyn,sleep until trie is built

	str = []rune("I want to know wether Marry is Chinese or Marry is American,and I think Chinese is great and American is great")
	filter.FilterText(dicFile1, str, []rune{}, '*')
	for _, char := range str {
		fmt.Printf("%c", char)
	}
	fmt.Printf("\n------------------------------------------------\n")
	//output: I want to know wether **************** or *****************,and I think **************** and *****************]

	str = []rune("I want to know wether Marry is[[ ]]Chinese or Marry is American,and I think Chinese is great and American is great")
	filter.FilterText(dicFile1, str, []rune{'[', ']'}, '*')
	for _, char := range str {
		fmt.Printf("%c", char)
	}
	fmt.Printf("\n------------------------------------------------\n")
	//output: I want to know wether ********[[*]]******* or *****************,and I think **************** and *****************
}
