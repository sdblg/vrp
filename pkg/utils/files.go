package utils

import (
	"bufio"
	"os"
)

// Reading txt file line by line and send to out channel
func ReadFile(path string, out chan string) {
	 // open input file
	 openedInputFile, err := os.Open(path)
	 PanicIfErr(err)
	 // close fi on exit and check for its returned error
	 defer func() {
		 PanicIfErr(openedInputFile.Close())
	 }()	 
 
	 // make a scanner for reading text file line by line
	 scanner := bufio.NewScanner(openedInputFile)	 
	 for scanner.Scan() {
		out <- scanner.Text()
	 }
	 close(out)
}