package main

import (
	"fmt"
	"os"
	"runtime"
)

func sayToFile(s string, f string) (bool, error) {
	fout, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return false, err
	} else {
		// fmt.Println("file opened.")
	}
	defer fout.Close()
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fout.WriteString(s)
		fout.WriteString("\n")
	}
	return true, nil
}

func createFile(p string, n string) (string, error) {
	sFileName := p + n
	_, e := os.Open(sFileName)
	if e == nil {
		fmt.Println("File already exists.")
		return sFileName, nil
	}
	e = os.MkdirAll(p, 0777)
	if e != nil {
		if os.IsPermission(e) {
			fmt.Println("No rights to perform such an operation...")
		}
		return "", e
	}
	_, e = os.Create(sFileName)
	if e != nil {
		fmt.Println("Unknown error, file not created.")
		return "", e
	} else {
		fmt.Println("File created.")
		return sFileName, nil
	}
}

func main() {
	var sFileName string
	_, e1 := fmt.Scanf("%s", &sFileName)
	if e1 != nil {
		fmt.Println(e1)
		fmt.Println("Using the default filename...")
		sFileName = "e:\\goSource\\testWords4.txt"
	}
	f, e := createFile("e:\\goSource\\", "testWords4.txt")
	if e != nil {
		fmt.Println(e)
		fmt.Println(f + " not created...")
	}
	sayToFile("world", f)
	go sayToFile("hello", f)

	//say("!!!!")
}
