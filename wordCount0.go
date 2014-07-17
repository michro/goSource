package main

import (
	"fmt"
	//	"unicode/utf8"
	"io/ioutil"
	"os"
	"time"
)

//main func
func main() {
	/*	fmt.Println(isLegal('e'))
		fmt.Println(isLegal('E'))
		fmt.Println(isLegal('_'))
		fmt.Println(isLegal(' '))
		fmt.Println(isLegal('\n'))*/
	fmt.Println("Please input the path of file to be counted...")
	var sFile string
	_, e1 := fmt.Scanf("%s", &sFile)
	if e1 != nil {
		fmt.Println(e1)
		fmt.Println("Using the default testWords.txt...")
		sFile = "e:\\goSource\\testWords2.txt"
	}
	fin, e2 := os.Open(sFile)
	defer fin.Close()
	if e2 != nil {
		fmt.Println(sFile, e2)
		return
	}
	f, e3 := ioutil.ReadAll(fin)
	if e3 != nil {
		panic(e3)
	} else {
		wm := countWordsInString(string(f))
		for k, v := range wm {
			fmt.Printf("Key: %s, Value: %d\n", k, v)
		}
	}
}

//count the appearance of a word
func countWordsInString(sInput string) map[string]int {
	//change the input into a rune array
	sTemp := []rune(sInput)
	return (countWordsInRunes(sTemp))
}

//count the appearance of a word version 2
func countWordsInRunes(sTemp []rune) map[string]int {
	//build a map for kv
	wordMap := make(map[string]int)

	//find a word
	i, lInput := 0, len(sTemp)
	fmt.Printf("length of input is %d\n", lInput)
	start := time.Now()
	fmt.Println("Time is ", start)
	for i < lInput {
		//make a slice for a word
		sWord := make([]rune, 100)
		j := 0
		for i < lInput && isLegal(sTemp[i]) {
			sWord[j] = sTemp[i]
			i++
			j++
		}

		//add a kv
		word := string(sWord[:j])
		fmt.Println(word)
		fmt.Printf("Start to find the word %s...\n", word)

		v, ok := wordMap[word]
		if ok {
			wordMap[word] = v + 1
			fmt.Println("found!")
		} else {
			wordMap[word] = 1
			fmt.Println("Not found...")
		}
		fmt.Println("Time is ", time.Now())
		//jump to next word
		i++
	}
	fmt.Printf("Cost %v\n", time.Now().Sub(start))
	return wordMap
}

//whether c is a legal word component
func isLegal(c rune) bool {
	switch {
	case 'a' <= c && c <= 'z':
		return true
	case 'A' <= c && c <= 'Z':
		return true
	case '_' == c:
		return true
	default:
		return false
	}
}
