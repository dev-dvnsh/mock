package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func generateName() string {
	randomInt1 := rand.IntN(50)
	randomInt2 := rand.IntN(50)
	name := FirstName[randomInt1] + " " + LastName[randomInt2]
	return name
}

var name = generateName()

func generateEmail() string {
	nameSlice := strings.Split(name, " ")
	var strNum strings.Builder
	for range 5 {
		num := rand.IntN(10)
		strNum.WriteString(strconv.Itoa(num))
	}

	email := strings.ToLower(nameSlice[0]) + string(strings.ToLower(nameSlice[1])[0]) + strNum.String() + "@gmail.com"
	return email
}

func generateAge() int {
	max := 65
	min := 18
	return rand.IntN(max-min) + min
}

// func generate(schemaField string) string{
// 	switch(schemaField){
// 	case "name":
// 	return
// 	}
// }

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

	fmt.Println(name)
	fmt.Println(generateEmail())
	fmt.Println(generateAge())
}
