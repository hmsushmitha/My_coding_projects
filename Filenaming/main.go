package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Name string

func main() {
	fmt.Println("Enter the file name")
	//Name := "birthday_001.txt"
	fmt.Scan(&Name)
	// => Birthday - 1 of 4.txt
	newName, err := match(Name, 4)
	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}
	fmt.Println(newName)
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(Name string, total int) (string, error) {
	// "birthday", "001", "txt"
	pieces := strings.Split(Name, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", Name)
	}
	// Birthday - 1.txt
	return fmt.Sprintf("%s - %d of %d.%s", (name), number, total, ext), nil
}
