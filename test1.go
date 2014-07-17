package main
import{
	"fmt"
	"unicode/utf8"
}

//print n times A

func printA(n int){
	a :=""
	for i:=0;i<n;i++{
		a+="A"
	}
	fmt.Println(a)
}

//count the appearance of a word 

func countWords(sInput string){
	wordMap := make(map[string]int)
	fmt.Println(wordMap)
}

func main(){
	n:=101
	for i:=1;i<n;i++{
		printA(i)
	}
	s := "abcabcabcbcbcbcc"
	fmt.Printf("String %s has length %d, runes %d.\n", s, len([]byte(s)), utf8.RuneCount([]byte(s)))
}
