package main

import (
	"fmt"
	"io/ioutil"
)

func readingDirrectories(path string, counter int) {
	streamDir, err := ioutil.ReadDir(path)
	if err != nil {
		panic("error opening homedir")
	}

	for _, files := range streamDir {
		for i := 0; i < counter; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s   (%d bytes)\n", files.Name(), files.Size())
		if files.IsDir() == true {

			readingDirrectories(path+"/"+files.Name(), counter+len(files.Name()))
		}

	}
}

func main() {

	readingDirrectories("/home/infiniss/dirwalk", 0)

}
