package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var numToWordMap = map[string]string{
	"0": "Zero",
	"1": "One",
	"2": "Two",
	"3": "Three",
	"4": "Four",
	"5": "Five",
	"6": "Six",
	"7": "Seven",
	"8": "Eight",
	"9": "Nine",
}

type Num struct {
	value string
}

// isNumeric function to check if string is numeric
func (n *Num) isNumeric() error {
	if n.value == "" {
		return errors.New("empty string is not a value")
	}
	_, err := strconv.ParseFloat(n.value, 64)
	if err != nil {
		return errors.New(n.value + " is not a number")
	}
	return nil
}

// ToWord function to convert string of numbers to words
func (n *Num) ToWord() (string, error) {
	err := n.isNumeric()
	if err != nil {
		return "", err
	}

	str := ""

	s := n.value

	if s[0] == '-' {
		str += "Negative"
		s = s[1:]
	}

	for _, ch := range s {
		str += numToWordMap[string(ch)]
	}
	return str, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No args were given")
		return
	}
	nums := os.Args[1:]

	for i := 0; i < len(nums); i++ {

		v := Num{nums[i]}
		str, err := v.ToWord()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(str)
		}

		if i < len(nums)-1 {
			fmt.Print(",")
		} else {
			// add a new line at the end of print
			fmt.Println()
		}
	}
}
