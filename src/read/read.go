package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type userInfo struct {
	fname string
	lname string
}

func main() {
	var filename string

	fmt.Println("give filename")
	_, err := fmt.Scan(&filename)

	if err != nil {
		fmt.Println("bad input")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("couldn't read file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var userInfoSlice []userInfo

	for scanner.Scan() {
		currentLine := scanner.Text()
		fullName := strings.Split(currentLine, " ")

		if len(fullName) != 2 {
			fmt.Println("line format sanity failed :", currentLine)
			continue
		}

		fname := fullName[0]
		lname := fullName[1]

		if len(fname) > 20 {
			fname = fname[:20]
		}

		if len(lname) > 20 {
			lname = lname[:20]
		}

		user := userInfo{fname: fname, lname: lname}
		userInfoSlice = append(userInfoSlice, user)
	}

	for _, elmt := range userInfoSlice {
		fmt.Println(elmt.fname, " ", elmt.lname)
	}
}
