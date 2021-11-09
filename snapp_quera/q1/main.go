package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	a := re.FindAllString("abc123def987asdf", -1)
	fmt.Println()
}

func makeNewNums(s string) []string {
	return []string
}

func Main2() {

	s := bufio.NewScanner(os.Stdin)

	var input string
	flag := true
	for flag {
		select {
		case <-time.After(time.Second * 1):
			flag = false
			break
		default:
			input += s.Text() + "\n"
		}
		fmt.Println("here")
	}
	fmt.Println(replaceNums(input))

}

func replaceNums(s string) string {
	s = strings.ReplaceAll(s, "1 ", "1st ")
	s = strings.ReplaceAll(s, "2 ", "2nd ")
	s = strings.ReplaceAll(s, "3 ", "3rd ")
	s = strings.ReplaceAll(s, "4 ", "4rd ")
	s = strings.ReplaceAll(s, "5 ", "5rd ")
	s = strings.ReplaceAll(s, "6 ", "6rd ")
	s = strings.ReplaceAll(s, "7 ", "7rd ")
	s = strings.ReplaceAll(s, "8 ", "8rd ")
	s = strings.ReplaceAll(s, "9 ", "9rd ")
	s = strings.ReplaceAll(s, "0 ", "0rd ")
	return s
}
