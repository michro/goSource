package main
import("fmt"
//	"unicode/utf8"
	"time"
	"os"
	"io/ioutil"
	//"strings"
)

//main func
func main(){
/*	fmt.Println(isLegal('e'))
	fmt.Println(isLegal('E'))
	fmt.Println(isLegal('_'))
	fmt.Println(isLegal(' '))
	fmt.Println(isLegal('\n'))*/
	fmt.Println("Please input the path of file to be counted...")
	var sFile string
	_,e1:=fmt.Scanf("%s\n", &sFile)
	if e1!=nil{
		fmt.Println(e1)
		fmt.Println("Using the default testWords.txt...")
		sFile = "e:\\goSource\\testWords3.txt"
	}
	fin,e2:= os.Open(sFile)
	defer fin.Close()
	if e2!=nil{
		fmt.Println(sFile,e2)
		return
	}
	f,e3 := ioutil.ReadAll(fin)
	if e3!=nil{
		panic(e3)
	}else{
		wm:=countWordsInString(string(f))
		fmt.Println("Length of the map",len(wm))
		printMap(wm)
	}
}

//print map
func printMap (m map[string]int) int{
	for k,v:=range m{
		fmt.Printf("Key: %q, Value: %d\n", k,v)
	}
	return 0
}

//count the appearance of a word 
func countWordsInString(sInput string) map[string]int{
	//change the input into a rune array
	sTemp := []rune(sInput)
	return(countWordsInRunes(sTemp))
}
//count the appearance of a word version 2 
func countWordsInRunes(sTemp []rune) map[string]int{
	//build a map for kv
	wordMap := make(map[string]int)
	
	//find a word
	i,lInput:=0,len(sTemp)
	//fmt.Printf("length of input is %d\n",lInput)
	start:=time.Now()
	//fmt.Println("Time is ", start)
	word := ""
	for i=0;i< lInput;i++{
		//make a slice for a word
		if isLegal(sTemp[i]){
			word=word+string(sTemp[i])
			//fmt.Printf("Current word %q...\n", word)
		}else{
			//word := string(sWord[:j])
			if len(word)>0{
				//fmt.Printf("The word is %q, with length %d\n", word, len(word))
				v,ok :=wordMap[word]
				if ok{
					wordMap[word]=v+1
					//fmt.Println("found!")
				}else{
					wordMap[word]=1
					//fmt.Println("Not found...")
				}
				word = ""
			}
		}
		
		//add a kv
//		fmt.Println(word)
//		fmt.Printf("Start to find the word %s...\n", word)
//		fmt.Println("Time is ", time.Now())

	}
	fmt.Printf("Cost %v\n",time.Now().Sub(start))
	return wordMap
}

//whether c is a legal word component
func isLegal(c rune) bool{
	switch{
		case 'a' <= c && c <= 'z':return true
		case 'A' <= c && c <= 'Z':return true
		case '_' == c :return true
		default: return false
	}
}

//sort a map
//not finished
func sortMap(m map[string]int) bool{
	for k, v := range m{
		fmt.Printf("Key: %q, Value: %d\n",k,v)
	}
	return true
}

