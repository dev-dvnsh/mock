package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// cmd := os.Args[0]
	// fmt.Println(cmd)
	if len(os.Args) <= 2 {
		fmt.Println("pass atleast 2 arguments after mock like mock name 10")
		os.Exit(0)
	}

	cmdArgs := os.Args[1]
	// fmt.Println(cmdArgs)
	schema := strings.Split(cmdArgs, " ")

	cmd1 := os.Args[2]

	num, err := strconv.Atoi(cmd1)
	if err != nil {
		log.Fatalf("Failed to convert: %v", err)
		os.Exit(0)
	}
	fmt.Println(schema)
	fmt.Println(num)
	for i := 3; i < len(os.Args); i++ {
		// fmt.Println(os.Args[i])
		// if os.Args[i][0:2] == "--" {
		// 	fmt.Println("found")
		// }
		if os.Args[i] == "--format" {
			fmt.Println("--found found")
			fileType := strings.TrimSpace(os.Args[i+1])

			fmt.Println("File Type: ", fileType)
		}
		if os.Args[i] == "--output" {
			fmt.Println("--output found")
			outputFileName := strings.TrimSpace(os.Args[i+1])

			fmt.Println("Output File Name: ", outputFileName)
		}
	}
}
