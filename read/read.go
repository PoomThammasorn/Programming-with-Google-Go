package main

import (
	"fmt"
	"os"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	var fileName string
	var usr []name

	fmt.Println("First, You MUST make a text file, using any text editor. And name it with .txt extension.")
	fmt.Println("Second, Enter the path of the file including the file name.")
	fmt.Print("e.g. /home/user/Documents/file.txt\n\n")
	fmt.Println("If eror occurs, ':no such file or directory'. It's not a problem of this program.")
	fmt.Println("It's because you entered the wrong path or the file doesn't exist.")
	fmt.Print("Best solution: Put my .go file in the same directory as your file.\n\n")
	fmt.Print("Enter the path of the file: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		var fname string
		var lname string
		_, err := fmt.Fscanln(file, &fname, &lname)
		if err != nil {
			break
		}
		usr = append(usr, name{fname, lname})
	}
	file.Close()
	for _, v := range usr {
		fmt.Println(v.fname, v.lname)
	}
}
