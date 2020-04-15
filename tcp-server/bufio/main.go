package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "some text \n that has nothing to do \n with your dreams"
	scanner := bufio.NewScanner(strings.NewReader(s))

	//scanner.Split(bufio.ScanWords) will split for each word

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
