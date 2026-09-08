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

func generateEmail(name string) string {
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

func generateUUID() string {
	randNum1 := rand.Uint32()
	randNum2 := rand.Uint32()
	randNum3 := rand.Uint32()
	randNum4 := rand.Uint32()
	hexString := fmt.Sprintf("%x%x%x%x", randNum1, randNum2, randNum3, randNum4)
	uuid := hexString[0:8] + "-" + hexString[8:12] + "-" + hexString[12:16] + "-" + hexString[16:20] + "-" + hexString[20:]

	return uuid
}

func generatePhone() string {
	var strPNum strings.Builder

	for range 10 {
		num := rand.IntN(10)
		strPNum.WriteString(strconv.Itoa(num))
	}
	phoneNumber := "+91-" + strPNum.String()
	return phoneNumber
}

func generateBool() bool {
	randNum := rand.IntN(2)
	if randNum == 0 {
		return false
	} else {
		return true
	}
}

func generateInt() int {
	randNum := rand.IntN(9000) + 1000
	return randNum
}

func generateDefault(field []string) string {
	fieldString := strings.Join(field, ", ")
	returnString := fieldString + " not recognised, use help for field names"
	return returnString
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(length int) string {
	var result strings.Builder
	for range length {
		index := rand.IntN(len(charset))
		result.WriteByte(charset[index])
	}
	return result.String()
}

func generate(schemaField string, name string) string {
	switch schemaField {
	case "name":
		return generateName()
	case "email":
		return generateEmail(name)
	case "age":
		return strconv.Itoa(generateAge())
	case "city":
		lastAddressIndex = rand.IntN(len(Addresses))
		return Addresses[lastAddressIndex].City
	case "country":
		return Addresses[lastAddressIndex].Country
	case "date":
		year := rand.IntN(2024-1980) + 1980
		month := rand.IntN(12) + 1
		day := rand.IntN(28) + 1
		return fmt.Sprintf("%d-%02d-%02d", year, month, day)
	case "uuid":
		return generateUUID()
	case "phone":
		return generatePhone()
	case "bool":
		return strconv.FormatBool(generateBool())
	case "int":
		return strconv.Itoa(generateInt())
	case "password":
		return generatePassword(8)
	default:
		return generateDefault([]string{schemaField})
	}
}

func main() {
	// cmd := os.Args[0]
	// fmt.Println(cmd)
	if len(os.Args) <= 2 {
		fmt.Println("Usage: mock \"<fields>\" <count> [--format json|csv] [--output filename]\nExample: mock \"name, email, age\" 100 --format csv")
		os.Exit(0)
	}

	cmdArgs := os.Args[1]
	// fmt.Println(cmdArgs)
	schema := strings.Split(cmdArgs, ",")
	for i := range schema {
		schema[i] = strings.TrimSpace(schema[i])
	}

	cmd1 := os.Args[2]

	num, err := strconv.Atoi(cmd1)
	if err != nil {
		log.Fatalf("Failed to convert: %v", err)
	}

	format := "json"
	outputFile := ""

	for i := 3; i < len(os.Args); i++ {

		if os.Args[i] == "--format" {
			if i+1 < len(os.Args) {
				format = strings.TrimSpace(os.Args[i+1])
			} else {
				fmt.Println("Error: --format requires a value")
				os.Exit(1)
			}
			fmt.Println("File Type: ", format)
		}
		if os.Args[i] == "--output" {

			if i+1 < len(os.Args) {
				outputFile = strings.TrimSpace(os.Args[i+1])
			} else {
				fmt.Println("Error: --output requires a value")
				os.Exit(1)
			}

			fmt.Println("Output File Name: ", outputFile)
		}

	}
	records := []map[string]string{}
	for range num {
		record := map[string]string{}
		for _, field := range schema {
			record[field] = generate(field, "")
		}
		records = append(records, record)
	}
	fmt.Println(records[0])
}
